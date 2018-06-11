package constants

// HOSTS
const MongoDBHost = "localhost"
const LogHost = "http://csd-elk:5000"
const DefaultServerPort = ":80"
const DefaultTimeout = 10000

// ROUTES
const PathLogin = "/login"
const PathRegister = "/register"
const PathAPI = "/api"
const PathCrud = "/data"

// MODELS
const AttrUserLogin = "login"
const AttrUserPass = "password"
const AttrUserID = "ID"
const AttrToken = "token"

// MESSAGES
const MessageStartingServer = "Starting server"
const MessageNewUserCreate = "New User "
const MessageNewUserLogin = "New login "
const MessageTryConnectDB = "Trying connect with mongodb"
const MessageConnectDBSuccess = "Connected with mongodb"

// ERRORS
const ErrorLoginAlreadyExists = "Login already exist!"
const ErrorInvalidUserOrPass = "Invalid User/Pass"
const ErrorInvalidToken = "Invalid Token"
const ErrorTryingConnectDB = "Error connecting with DB"
const ErrorRegisterNewUser = "Error when try register new User"
const ErrorLogin = "Error when trying login"
const ErrorSenddingLog = "Cannot send log to server"
