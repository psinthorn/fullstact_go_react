module.exports = {
  mongodb: {
    server: 'mongo',
    port: 27017,
    admin: true,
    auth: [
      {
        database: 'logs',
        username: 'admin',
        password: 'admin',
      },
    ],
  },
  site: {
    baseUrl: '/',
    cookieKeyName: 'mongo-express',
    cookieSecret: 'cookiesecret',
    host: '0.0.0.0',
    port: 8081,
    requestSizeLimit: '50mb',
    sessionSecret: 'sessionsecret',
    sslEnabled: false,
    sslCert: '',
    sslKey: '',
  },
};