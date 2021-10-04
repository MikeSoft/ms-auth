import tornado.ioloop
import tornado.web
import signal
from firebase_admin import auth
import firebase_admin
from firebase_admin import credentials

json_route = "adminsdk.json"
cred = credentials.Certificate(json_route)
f_app=firebase_admin.initialize_app(cred, name='ms-authentication')

class MainHandler(tornado.web.RequestHandler):

    def get(self):
        try:
            id_token = self.request.headers.get('Authorization', '')
            decoded_token = auth.verify_id_token(id_token,app=f_app)
            uid = decoded_token['uid']
            self.set_status(200)
            self.set_header("UID", uid)
        except Exception as ex:
            self.clear()
            self.set_status(401)    

def make_app():
    return tornado.web.Application([
        (r"/auth/request", MainHandler),
    ])

if __name__ == "__main__":
    app = make_app()
    app.listen(8080)
    print("Server initialized")
    tornado.ioloop.IOLoop.instance().start()