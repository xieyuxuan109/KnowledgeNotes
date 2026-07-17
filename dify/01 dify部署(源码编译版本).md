## 克隆仓库
```
git clone https://github.com/langgenius/dify.git
cd dify/docker
cp envs/middleware.env.example middleware.env
docker compose --env-file middleware.env -f docker-compose.middleware.yaml -p dify up -d
```
## 设置后端服务
```
cd ../api
cp .env.example .env
awk -v key="$(openssl rand -base64 42)" '/^SECRET_KEY=/ {sub(/=.*/, "=" key)} 1' .env > temp_env && mv temp_env .env
snap install astral-uv --classic
uv sync --dev
uv run flask db upgrade
uv run flask run --host 0.0.0.0 --port=5001 --debug
```
## 预期输出
```
* Debug mode: on
INFO:werkzeug:WARNING: This is a development server. Do not use it in a production deployment. Use a production WSGI server instead.
 * Running on all addresses (0.0.0.0)
 * Running on http://127.0.0.1:5001
INFO:werkzeug:Press CTRL+C to quit
INFO:werkzeug: * Restarting with stat
WARNING:werkzeug: * Debugger is active!
INFO:werkzeug: * Debugger PIN: 695-801-919
```
## 启动worker服务
```
uv run celery -A app.celery worker -P gevent -c 1 --loglevel INFO -Q dataset,dataset_summary,priority_dataset,priority_pipeline,pipeline,mail,ops_trace,app_deletion,plugin,workflow_storage,conversation,workflow,schedule_poller,schedule_executor,triggered_workflow_dispatcher,trigger_refresh_executor,retention,workflow_based_app_execution
```
## 预期输出
```
-------------- celery@bwdeMacBook-Pro-2.local v5.4.0 (opalescent)
--- ***** -----
-- ******* ---- macOS-15.4.1-arm64-arm-64bit 2025-04-28 17:07:14
- *** --- * ---
- ** ---------- [config]
- ** ---------- .> app:         app_factory:0x1439e8590
- ** ---------- .> transport:   redis://:**@localhost:6379/1
- ** ---------- .> results:     postgresql://postgres:**@localhost:5432/dify
- *** --- * --- .> concurrency: 1 (gevent)
  -- ******* ---- .> task events: OFF (enable -E to monitor tasks in this worker)
  --- ***** -----
  -------------- [queues]
  .> dataset          exchange=dataset(direct) key=dataset
  .> generation       exchange=generation(direct) key=generation
  .> mail             exchange=mail(direct) key=mail
  .> ops_trace        exchange=ops_trace(direct) key=ops_trace

[tasks]
. schedule.clean_embedding_cache_task.clean_embedding_cache_task
. schedule.clean_messages.clean_messages
. schedule.clean_unused_datasets_task.clean_unused_datasets_task
. schedule.create_tidb_serverless_task.create_tidb_serverless_task
. schedule.mail_clean_document_notify_task.mail_clean_document_notify_task
. schedule.update_tidb_serverless_status_task.update_tidb_serverless_status_task
. tasks.add_document_to_index_task.add_document_to_index_task
. tasks.annotation.add_annotation_to_index_task.add_annotation_to_index_task
. tasks.annotation.batch_import_annotations_task.batch_import_annotations_task
. tasks.annotation.delete_annotation_index_task.delete_annotation_index_task
. tasks.annotation.disable_annotation_reply_task.disable_annotation_reply_task
. tasks.annotation.enable_annotation_reply_task.enable_annotation_reply_task
. tasks.annotation.update_annotation_to_index_task.update_annotation_to_index_task
. tasks.batch_clean_document_task.batch_clean_document_task
. tasks.batch_create_segment_to_index_task.batch_create_segment_to_index_task
. tasks.clean_dataset_task.clean_dataset_task
. tasks.clean_document_task.clean_document_task
. tasks.clean_notion_document_task.clean_notion_document_task
. tasks.deal_dataset_vector_index_task.deal_dataset_vector_index_task
. tasks.delete_account_task.delete_account_task
. tasks.delete_segment_from_index_task.delete_segment_from_index_task
. tasks.disable_segment_from_index_task.disable_segment_from_index_task
. tasks.disable_segments_from_index_task.disable_segments_from_index_task
. tasks.document_indexing_sync_task.document_indexing_sync_task
. tasks.document_indexing_task.document_indexing_task
. tasks.document_indexing_update_task.document_indexing_update_task
. tasks.duplicate_document_indexing_task.duplicate_document_indexing_task
. tasks.enable_segments_to_index_task.enable_segments_to_index_task
. tasks.mail_account_deletion_task.send_account_deletion_verification_code
. tasks.mail_account_deletion_task.send_deletion_success_task
. tasks.mail_email_code_login.send_email_code_login_mail_task
. tasks.mail_invite_member_task.send_invite_member_mail_task
. tasks.mail_reset_password_task.send_reset_password_mail_task
. tasks.ops_trace_task.process_trace_tasks
. tasks.recover_document_indexing_task.recover_document_indexing_task
. tasks.remove_app_and_related_data_task.remove_app_and_related_data_task
. tasks.remove_document_from_index_task.remove_document_from_index_task
. tasks.retry_document_indexing_task.retry_document_indexing_task
. tasks.sync_website_document_indexing_task.sync_website_document_indexing_task

2025-04-28 17:07:14,681 INFO [connection.py:22]  Connected to redis://:**@localhost:6379/1
2025-04-28 17:07:14,684 INFO [mingle.py:40]  mingle: searching for neighbors
2025-04-28 17:07:15,704 INFO [mingle.py:49]  mingle: all alone
2025-04-28 17:07:15,733 INFO [worker.py:175]  celery@bwdeMacBook-Pro-2.local ready.
2025-04-28 17:07:15,742 INFO [pidbox.py:111]  pidbox: Connected to redis://:**@localhost:6379/1.
```
## 启动beat服务
```
uv run celery -A app.celery beat
```
## 设置web服务
```
apt install npm -y //注意版本 要比较新的
npm i -g pnpm
cd ../web
pnpm install --frozen-lockfile
```
## 构建web服务
```
准备环境变量配置文件
在当前目录中创建一个名为 .env.local 的文件，并从 .env.example 复制内容。根据你的需求修改这些环境变量的值：
# For production release, change this to PRODUCTION
NEXT_PUBLIC_DEPLOY_ENV=DEVELOPMENT

# The deployment edition, SELF_HOSTED or CLOUD
NEXT_PUBLIC_EDITION=SELF_HOSTED

# The base URL of console application, refers to the Console base URL of WEB service if console domain is different from api or web app domain.
# example: http://cloud.dify.ai/console/api
NEXT_PUBLIC_API_PREFIX=http://localhost:5001/console/api

# The URL for Web APP, refers to the Web App base URL of WEB service if web app domain is different from console or api domain.
# example: http://udify.app/api
NEXT_PUBLIC_PUBLIC_API_PREFIX=http://localhost:5001/api

# When the frontend and backend run on different subdomains, set NEXT_PUBLIC_COOKIE_DOMAIN=1.
NEXT_PUBLIC_COOKIE_DOMAIN=

# SENTRY
NEXT_PUBLIC_SENTRY_DSN=
NEXT_PUBLIC_SENTRY_ORG=
NEXT_PUBLIC_SENTRY_PROJECT=
```