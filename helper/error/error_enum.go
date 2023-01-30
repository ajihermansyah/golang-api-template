package error

const (
	FailUnmarshalResponseBodyError ErrorType = "ER1000 Fail to unmarshal response body"   // used by controller
	FailUnmarshalRequestBodyError  ErrorType = "ER1000 Fail to unmarshal request body"    // used by controller
	FailUnmarshalTokenError        ErrorType = "ER1000 Fail to unmarshal token into user" // used by controller
	AnyInvalidParameters           ErrorType = "ER1000 there are some invalid parameters" // used by controller
	ApiKeyIsNeeded                 ErrorType = "ER1000 Api Key is needed on the header"   // used by controller
	EntityNotFound                 ErrorType = "ER1001 Entity %s with id %s is not found" // used by injected repo in interactor
	UnrecognizedEnum               ErrorType = "ER1002 %s is not recognized %s enum"      // used by enum
	DatabaseNotFoundInContextError ErrorType = "ER1003 Database is not found in context"  // used by repoimpl
)
const EnvironmentMustNotEmpty ErrorType = "ER1000 environment must not empty"  //
const CredFileNameMustNotEmpty ErrorType = "ER1000 cred must not empty"        //
const BucketMustNotEmpty ErrorType = "ER1000 bucket must not empty"            //
const DirectoryMustNotEmpty ErrorType = "ER1000 directory must not empty"      //
const RandomIDMust4CharLength ErrorType = "ER1000 random id must4 char length" //
const InvalidDate ErrorType = "ER1000 invalid date"                            //
const SomethingHappen ErrorType = "ER1000 something happen"                    //

const UsernameMustNotEmpty ErrorType = "ER1000 username must not empty"        //
const FullnameMustNotEmpty ErrorType = "ER1001 fullname must not empty"        //
const PasswordMustNotEmpty ErrorType = "ER1002 password must not empty"        //
const OldPasswordMustNotEmpty ErrorType = "ER1003 old password must not empty" //
const NewPasswordMustNotEmpty ErrorType = "ER1004 new password must not empty" //
const RoleMustNotEmpty ErrorType = "ER1005 role must not empty"                //
const CompanyMustNotEmpty ErrorType = "ER1005 company must not empty"          //
const CourierTypeMustNotEmpty ErrorType = "ER1005 courier type must not empty" //
const UsernameHasTaken ErrorType = "ER1006 username has been taken"            //
const AgeMustNotEmpty ErrorType = "ER1007 age must not empty"
const GenderMustNotEmpty ErrorType = "ER1007 gender must not empty"
const EmailMustNotEmpty ErrorType = "ER1007 email must not empty"

const DuplicateRole ErrorType = "ER1007 cannot create duplicate role"                 //
const RoleHasUsed ErrorType = "ER1008 please make sure no users were using this role" //

const InsufficientRole ErrorType = "ER1009 make sure your role has sufficient authorities" //

const NoUserLogged ErrorType = "ER1010 no username extracted while logging"                                   //
const NoRoleLogged ErrorType = "ER1011 no role extracted while logging"                                       //
const LogJsonParams ErrorType = "ER1012 error while converting json-formatted params to string while logging" //

const AppNameMustNotEmpty ErrorType = "ER1000 client name must not empty"    //
const ClientNameMustNotEmpty ErrorType = "ER1000 client name must not empty" //
const VersionMustNotEmpty ErrorType = "ER1000 version must not empty"        //
const IntentionMustNotEmpty ErrorType = "ER1000 intention must not empty"
const AuthTokenMustNotEmpty ErrorType = "ER1000 auth token must not empty"
const ServiceNameMustNotEmpty ErrorType = "ER1000 service name must not empty" //

const AppNameMustBeSent ErrorType = "ER1000 client name must be sent"    //
const ClientNameMustBeSent ErrorType = "ER1000 client name must be sent" //
const VersionMustBeSent ErrorType = "ER1000 version must be sent"        //
const IntentionMustBeSent ErrorType = "ER1000 intention must be sent"
