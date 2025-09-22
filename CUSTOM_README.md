### 内部API接口 (/api/)
机器人管理 (/api/bot)
├── POST /api/bot/get_type_list
└── POST /api/bot/upload_file

文件上传 (/api/common/upload)
├── GET /api/common/upload/apply_upload_action
├── POST /api/common/upload/apply_upload_action
└── POST /api/common/upload/*tos_uri

对话管理 (/api/conversation)
├── POST /api/conversation/break_message
├── POST /api/conversation/chat
├── POST /api/conversation/clear_message
├── POST /api/conversation/create_section
├── POST /api/conversation/delete_message
└── POST /api/conversation/get_message_list

开发者工具 (/api/developer)
└── POST /api/developer/get_icon

草稿机器人 (/api/draftbot)
├── POST /api/draftbot/commit_check
├── POST /api/draftbot/create
├── POST /api/draftbot/delete
├── POST /api/draftbot/duplicate
├── POST /api/draftbot/get_display_info
├── POST /api/draftbot/list_draft_history
├── POST /api/draftbot/publish
├── POST /api/draftbot/update_display_info
└── POST /api/draftbot/publish/connector/list

智能体项目 (/api/intelligence_api)
├── POST /api/intelligence_api/draft_project/copy
├── POST /api/intelligence_api/draft_project/create
├── POST /api/intelligence_api/draft_project/delete
├── POST /api/intelligence_api/draft_project/inner_task_list
├── POST /api/intelligence_api/draft_project/update
├── POST /api/intelligence_api/publish/check_version_number
├── POST /api/intelligence_api/publish/connector_list
├── POST /api/intelligence_api/publish/get_published_connector
├── POST /api/intelligence_api/publish/publish_project
├── POST /api/intelligence_api/publish/publish_record_detail
├── POST /api/intelligence_api/publish/publish_record_list
├── POST /api/intelligence_api/search/get_draft_intelligence_info
├── POST /api/intelligence_api/search/get_draft_intelligence_list
└── POST /api/intelligence_api/search/get_recently_edit_intelligence

知识库管理 (/api/knowledge)
├── POST /api/knowledge/create
├── POST /api/knowledge/delete
├── POST /api/knowledge/detail
├── POST /api/knowledge/list
├── POST /api/knowledge/update
├── POST /api/knowledge/document/create
├── POST /api/knowledge/document/delete
├── POST /api/knowledge/document/list
├── POST /api/knowledge/document/resegment
├── POST /api/knowledge/document/update
├── POST /api/knowledge/document/progress/get
├── POST /api/knowledge/icon/get
├── POST /api/knowledge/photo/caption
├── POST /api/knowledge/photo/detail
├── POST /api/knowledge/photo/extract_caption
├── POST /api/knowledge/photo/list
├── POST /api/knowledge/review/create
├── POST /api/knowledge/review/mget
├── POST /api/knowledge/review/save
├── POST /api/knowledge/slice/create
├── POST /api/knowledge/slice/delete
├── POST /api/knowledge/slice/list
├── POST /api/knowledge/slice/update
├── POST /api/knowledge/table_schema/get
└── POST /api/knowledge/table_schema/validate

商店管理 (/api/marketplace)
├── GET /api/marketplace/product/detail
├── GET /api/marketplace/product/list
├── POST /api/marketplace/product/duplicate
├── POST /api/marketplace/product/favorite
└── GET /api/marketplace/product/favorite/list.v2

记忆/数据库 (/api/memory)
├── GET /api/memory/doc_table_info
├── GET /api/memory/sys_variable_conf
├── GET /api/memory/table_mode_config
├── POST /api/memory/database/add
├── POST /api/memory/database/bind_to_bot
├── POST /api/memory/database/delete
├── POST /api/memory/database/get_by_id
├── POST /api/memory/database/get_connector_name
├── POST /api/memory/database/get_online_database_id
├── POST /api/memory/database/get_template
├── POST /api/memory/database/list
├── POST /api/memory/database/list_records
├── POST /api/memory/database/unbind_to_bot
├── POST /api/memory/database/update
├── POST /api/memory/database/update_bot_switch
├── POST /api/memory/database/update_records
├── POST /api/memory/database/table/list_new
├── POST /api/memory/database/table/reset
├── GET /api/memory/project/variable/meta_list
├── POST /api/memory/project/variable/meta_update
├── POST /api/memory/table_file/get_progress
├── POST /api/memory/table_file/submit
├── POST /api/memory/table_schema/get
├── POST /api/memory/table_schema/validate
├── POST /api/memory/variable/delete
├── POST /api/memory/variable/get
├── POST /api/memory/variable/get_meta
└── POST /api/memory/variable/upsert

OAuth认证 (/api/oauth)
└── GET /api/oauth/authorization_code

用户认证 (/api/passport)
├── POST /api/passport/account/info/v2/
├── POST /api/passport/web/email/login/
├── GET /api/passport/web/email/password/reset/
├── POST /api/passport/web/email/register/v2/
└── GET /api/passport/web/logout/

权限管理 (/api/permission_api)
├── POST /api/permission_api/coze_web_app/impersonate_coze_user
├── POST /api/permission_api/pat/create_personal_access_token_and_permission
├── POST /api/permission_api/pat/delete_personal_access_token_and_permission
├── GET /api/permission_api/pat/get_personal_access_token_and_permission
├── GET /api/permission_api/pat/list_personal_access_tokens
└── POST /api/permission_api/pat/update_personal_access_token_and_permission

游乐场 (/api/playground)
├── POST /api/playground/get_onboarding
└── POST /api/playground/upload/auth_token

游乐场API (/api/playground_api)
├── POST /api/playground_api/create_update_shortcut_command
├── POST /api/playground_api/delete_prompt_resource
├── POST /api/playground_api/get_file_list
├── POST /api/playground_api/get_imagex_url
├── POST /api/playground_api/get_official_prompt_list
├── GET /api/playground_api/get_prompt_resource_info
├── POST /api/playground_api/mget_user_info
├── POST /api/playground_api/report_user_behavior
├── POST /api/playground_api/upsert_prompt_resource
├── POST /api/playground_api/draftbot/get_draft_bot_info
├── POST /api/playground_api/draftbot/update_draft_bot_info
├── POST /api/playground_api/operate/get_bot_popup_info
├── POST /api/playground_api/operate/update_bot_popup_info
└── POST /api/playground_api/space/list

插件管理 (/api/plugin)
└── POST /api/plugin/get_oauth_schema

插件开发 (/api/plugin_api)
├── POST /api/plugin_api/batch_create_api
├── POST /api/plugin_api/check_and_lock_plugin_edit
├── POST /api/plugin_api/convert_to_openapi
├── POST /api/plugin_api/create_api
├── POST /api/plugin_api/debug_api
├── POST /api/plugin_api/del_plugin
├── POST /api/plugin_api/delete_api
├── POST /api/plugin_api/get_bot_default_params
├── POST /api/plugin_api/get_dev_plugin_list
├── POST /api/plugin_api/get_oauth_schema
├── POST /api/plugin_api/get_oauth_status
├── POST /api/plugin_api/get_playground_plugin_list
├── POST /api/plugin_api/get_plugin_apis
├── POST /api/plugin_api/get_plugin_info
├── POST /api/plugin_api/get_plugin_next_version
├── POST /api/plugin_api/get_queried_oauth_plugins
├── POST /api/plugin_api/get_updated_apis
├── POST /api/plugin_api/get_user_authority
├── POST /api/plugin_api/library_resource_list
├── POST /api/plugin_api/project_resource_list
├── POST /api/plugin_api/publish_plugin
├── POST /api/plugin_api/register
├── POST /api/plugin_api/register_plugin_meta
├── POST /api/plugin_api/resource_copy_cancel
├── POST /api/plugin_api/resource_copy_detail
├── POST /api/plugin_api/resource_copy_dispatch
├── POST /api/plugin_api/resource_copy_retry
├── POST /api/plugin_api/revoke_auth_token
├── POST /api/plugin_api/unlock_plugin_edit
├── POST /api/plugin_api/update
├── POST /api/plugin_api/update_api
├── POST /api/plugin_api/update_bot_default_params
└── POST /api/plugin_api/update_plugin_meta

用户管理 (/api/user)
├── POST /api/user/update_profile
└── POST /api/user/update_profile_check

Web用户 (/api/web)
└── POST /api/web/user/update/upload_avatar/

工作流管理 (/api/workflow_api)
├── GET /api/workflow_api/apiDetail
├── POST /api/workflow_api/batch_delete
├── POST /api/workflow_api/cancel
├── POST /api/workflow_api/canvas
├── POST /api/workflow_api/copy
├── POST /api/workflow_api/copy_wk_template
├── POST /api/workflow_api/create
├── POST /api/workflow_api/delete
├── POST /api/workflow_api/delete_strategy
├── POST /api/workflow_api/example_workflow_list
├── GET /api/workflow_api/get_node_execute_history
├── GET /api/workflow_api/get_process
├── POST /api/workflow_api/get_trace
├── POST /api/workflow_api/history_schema
├── POST /api/workflow_api/list_publish_workflow
├── POST /api/workflow_api/list_spans
├── POST /api/workflow_api/llm_fc_setting_detail
├── POST /api/workflow_api/llm_fc_setting_merged
├── POST /api/workflow_api/nodeDebug
├── POST /api/workflow_api/node_panel_search
├── POST /api/workflow_api/node_template_list
├── POST /api/workflow_api/node_type
├── POST /api/workflow_api/publish
├── POST /api/workflow_api/released_workflows
├── POST /api/workflow_api/save
├── POST /api/workflow_api/sign_image_url
├── POST /api/workflow_api/test_resume
├── POST /api/workflow_api/test_run
├── POST /api/workflow_api/update_meta
├── POST /api/workflow_api/validate_tree
├── POST /api/workflow_api/workflow_detail
├── POST /api/workflow_api/workflow_detail_info
├── POST /api/workflow_api/workflow_list
├── POST /api/workflow_api/workflow_references
├── POST /api/workflow_api/chat_flow_role/create
├── POST /api/workflow_api/chat_flow_role/delete
├── GET /api/workflow_api/chat_flow_role/get
├── POST /api/workflow_api/project_conversation/create
├── POST /api/workflow_api/project_conversation/delete
├── GET /api/workflow_api/project_conversation/list
├── POST /api/workflow_api/project_conversation/update
└── POST /api/workflow_api/upload/auth_token

### OpenAPI v1接口 (/v1/)
对话管理
├── GET /v1/conversations
├── DELETE /v1/conversations/:conversation_id
├── PUT /v1/conversations/:conversation_id
├── POST /v1/conversations/:conversation_id/clear
├── POST /v1/conversation/create
└── POST /v1/conversation/message/list

应用管理
└── GET /v1/apps/:app_id

机器人管理
├── GET /v1/bot/get_online_info
└── GET /v1/bots/:bot_id

文件管理
└── POST /v1/files/upload

工作流管理
├── GET /v1/workflow/get_run_history
├── POST /v1/workflow/run
├── POST /v1/workflow/stream_resume
├── POST /v1/workflow/stream_run
├── POST /v1/workflow/conversation/create
├── POST /v1/workflows/chat
└── GET /v1/workflows/:workflow_id

OpenAPI v3接口 (/v3/)
聊天服务
├── POST /v3/chat
└── POST /v3/chat/cancel

### 接口统计
总接口数量: 218个
├── 内部API接口 (/api/*): 200+个
├── OpenAPI v1接口 (/v1/*): 15个
└── OpenAPI v3接口 (/v3/*): 2个

服务地址: http://localhost:8888/


### 创建新用户
```shell
curl -X POST http://localhost:8888/api/passport/web/email/register/v2/ \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@opencoze.com",
    "password": "admin"
  }'
```

### Configure Model
- [Model Configuration](https://github.com/coze-dev/coze-studio/wiki/3.-Model-configuration)
- Restart Server: `docker compose --profile '*' up -d --force-recreate --no-deps coze-server`


### Configure Plugin
- [Plugin Configuration](https://github.com/coze-dev/coze-studio/wiki/4.-Plugin-Configuration)

### 二开部分
