package constants

// HOSTS
const DefaultServerPort = ":80"
const DefaultTimeout = 10000

// TODO: Get by System Var
const MongoDBHost = "csd-gateway-db"
const LogHost = "http://logstash:5000"
const CrudHost = "http://csd-crud"

// ROUTES
const PathLogin = "/login"
const PathRegister = "/register"
const PathAPI = "/csd"
const PathCrud = "/api"

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
const MessageIsRequestingTo = " is requesting to "
const MessageSecretCreatingFile = "Creating secret file "

// ERRORS
const ErrorLoginAlreadyExists = "Login already exist!"
const ErrorInvalidUserOrPass = "Invalid User/Pass"
const ErrorInvalidToken = "Invalid Token"
const ErrorTryingConnectDB = "Error connecting with DB"
const ErrorRegisterNewUser = "Error when try register new User"
const ErrorLogin = "Error when trying login"
const ErrorSenddingLog = "Cannot send log to server"
const ErrorRequesting = "error on request "
const ErrorCrud = "Crud internal error "
const ErrorFileSecret = "Error creating secret file "
const ErrorDecodingSecret = "Error while try decode sercret "

// SECRET
const SecretFileName = "secret.txt"
const SecretSize = 32
