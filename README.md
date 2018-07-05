echo-session

copy from <https://github.com/astaxie/beego/tree/master/session>

session
==============

session is a Go session manager. It can use many session providers. Just like the `database/sql` and `database/sql/driver`.

## How to install?

	go get github.com/gocommon/echo-session


## What providers are supported?

As of now this session manager support memory, file, Redis and MySQL.


## How to use it?

		e.Use(session.Middleware(session.MiddlewareConfig{
			ManagerConfig: &session.ManagerConfig{Provider: "file", ProviderConfig: "./tmp"},
		}))

## How to write own provider?

When you develop a web app, maybe you want to write own provider because you must meet the requirements.

Writing a provider is easy. You only need to define two struct types 
(Session and Provider), which satisfy the interface definition. 
Maybe you will find the **memory** provider is a good example.

	type SessionStore interface {
		Set(key, value interface{}) error     //set session value
		Get(key interface{}) interface{}      //get session value
		Delete(key interface{}) error         //delete session value
		SessionID() string                    //back current sessionID
		SessionRelease(w *echo.Response) // release the resource & save data to provider & return the data
		Flush() error                         //delete all data
	}
	
	type Provider interface {
		SessionInit(gclifetime int64, config string) error
		SessionRead(sid string) (SessionStore, error)
		SessionExist(sid string) bool
		SessionRegenerate(oldsid, sid string) (SessionStore, error)
		SessionDestroy(sid string) error
		SessionAll() int //get all active session
		SessionGC()
	}


## LICENSE

BSD License http://creativecommons.org/licenses/BSD/
