redis_connection = Proc.new do
  Redis.new(
      host: ENV.fetch('REDIS_HOST') { 'redis' },
      port: ENV.fetch('REDIS_PORT') { 6379 },
      db:   ENV.fetch('REDIS_DB_SIDEKIQ') { 1 },
  )
end

Sidekiq.configure_server do |config|
  config.redis = ConnectionPool.new(size: ENV.fetch('RAILS_MAX_THREADS') { 5 }, &redis_connection)
end

Sidekiq.configure_client do |config|
  server_concurrency = Integer(ENV.fetch('SIDEKIQ_CONCURRENCY') { ENV.fetch('RAILS_MAX_THREADS') { 5 } })
  server_connections = server_concurrency + 2 # См. https://github.com/mperham/sidekiq/issues/1941#issuecomment-55137816
  config.redis = ConnectionPool.new(size: server_connections, &redis_connection)
end
