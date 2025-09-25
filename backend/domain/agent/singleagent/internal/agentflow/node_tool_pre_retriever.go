/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package agentflow

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/cloudwego/eino/compose"
	"github.com/coze-dev/coze-studio/backend/pkg/lang/conv"

	"github.com/coze-dev/coze-studio/backend/api/model/crossdomain/agentrun"
	workflowModel "github.com/coze-dev/coze-studio/backend/api/model/crossdomain/workflow"
	crossplugin "github.com/coze-dev/coze-studio/backend/crossdomain/contract/plugin"
	"github.com/coze-dev/coze-studio/backend/crossdomain/contract/plugin/consts"
	"github.com/coze-dev/coze-studio/backend/crossdomain/contract/plugin/model"
	crossworkflow "github.com/coze-dev/coze-studio/backend/crossdomain/contract/workflow"
	"github.com/coze-dev/coze-studio/backend/pkg/lang/ptr"
	"github.com/coze-dev/coze-studio/backend/pkg/logs"
)

type toolPreCallConf struct{}

func newPreToolRetriever(conf *toolPreCallConf) *toolPreCallConf {
	return &toolPreCallConf{}
}

func (pr *toolPreCallConf) toolPreRetrieve(ctx context.Context, ar *AgentRequest) ([]*schema.Message, error) {
	if len(ar.PreCallTools) == 0 {
		return nil, nil
	}

	var tms []*schema.Message
	for _, item := range ar.PreCallTools {

		var toolResp string
		switch item.Type {
		case agentrun.ToolTypePlugin:

			etr := &model.ExecuteToolRequest{
				UserID:          ar.UserID,
				ExecDraftTool:   false,
				PluginID:        item.PluginID,
				ToolID:          item.ToolID,
				ArgumentsInJson: item.Arguments,
				ExecScene: func(isDraft bool) consts.ExecuteScene {
					if isDraft {
						return consts.ExecSceneOfDraftAgent
					} else {
						return consts.ExecSceneOfOnlineAgent
					}
				}(ar.Identity.IsDraft),
			}

			opts := []model.ExecuteToolOpt{
				model.WithInvalidRespProcessStrategy(consts.InvalidResponseProcessStrategyOfReturnDefault),
				model.WithProjectInfo(&model.ProjectInfo{
					ProjectID:      ar.Identity.AgentID,
					ProjectType:    consts.ProjectTypeOfAgent,
					ProjectVersion: ptr.Of(ar.Identity.Version),
				}),
			}
			execResp, err := crossplugin.DefaultSVC().ExecuteTool(ctx, etr, opts...)
			logs.CtxInfof(ctx, "tool pre call plugin resp: %v, err: %v", conv.DebugJsonToStr(execResp), err)
			
			if err != nil {
				logs.CtxErrorf(ctx, "Failed to call ToolTypePlugin nodes: %v", err)
				return nil, err
			}
			toolResp = execResp.TrimmedResp

		case agentrun.ToolTypeWorkflow:
			var input map[string]any
			err := json.Unmarshal([]byte(item.Arguments), &input)
			if err != nil {
				logs.CtxErrorf(ctx, "Failed to unmarshal json arguments: %s", item.Arguments)
				return nil, err
			}
			execResp, _, err := crossworkflow.DefaultSVC().SyncExecuteWorkflow(ctx, workflowModel.ExecuteConfig{
				ID:           item.PluginID,
				ConnectorID:  ar.Identity.ConnectorID,
				ConnectorUID: ar.UserID,
				TaskType:     crossworkflow.TaskTypeForeground,
				AgentID:      ptr.Of(ar.Identity.AgentID),
				Mode: func() crossworkflow.ExecuteMode {
					if ar.Identity.IsDraft {
						return crossworkflow.ExecuteModeDebug
					} else {
						return crossworkflow.ExecuteModeRelease
					}
				}(),
			}, input)

			logs.CtxInfof(ctx, "tool pre call workflow resp: %v, err: %v", conv.DebugJsonToStr(execResp), err)
			
			if err != nil {
				// 检查是否是中断错误 - 这种情况说明工作流需要交互但我们没有提前检测到
				if pr.isWorkflowInterruptError(err) {
					logs.CtxInfof(ctx, "Workflow %d interrupted during execution (contains undetected interactive nodes), skipping", item.PluginID)
					continue
				}
				return nil, err
			}
			toolResp = ptr.From(execResp.Output)
		}

		if toolResp != "" {
			uID := uuid.New()
			// 去掉UUID中的破折号以符合OpenAI API 40字符限制 (call_ + 32字符UUID = 37字符)
			toolCallID := "call_" + strings.ReplaceAll(uID.String(), "-", "")
			tms = append(tms, &schema.Message{
				Role: schema.Assistant,
				ToolCalls: []schema.ToolCall{
					{
						Type: "function",
						Function: schema.FunctionCall{
							Name:      item.ToolName,
							Arguments: item.Arguments,
						},
						ID: toolCallID,
					},
				},
			})

			tms = append(tms, &schema.Message{
				Role:       schema.Tool,
				Content:    toolResp,
				ToolCallID: toolCallID,
			})
		}
	}

	return tms, nil
}

// isWorkflowInterruptError 检查错误是否为工作流中断错误
func (pr *toolPreCallConf) isWorkflowInterruptError(err error) bool {
	if err == nil {
		return false
	}
	
	// 检查是否是 Eino 的中断错误
	_, isInterruptErr := compose.ExtractInterruptInfo(err)
	if isInterruptErr {
		return true
	}
	
	// 检查错误消息中是否包含中断相关的关键词
	errMsg := err.Error()
	interruptKeywords := []string{
		"interrupt",
		"InterruptEvent",
		"interrupt and rerun",
		"NewInterruptAndRerunErr",
		"InputReceiver", // 特别检查 InputReceiver 相关的错误
		"QuestionAnswer", // 特别检查 QuestionAnswer 相关的错误
	}
	
	for _, keyword := range interruptKeywords {
		if strings.Contains(errMsg, keyword) {
			return true
		}
	}
	
	return false
}
