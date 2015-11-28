from flask import Flask
from fetch import fetch
from settings import PORT

app = Flask(__name__)

@app.route('/test')
def test():
    return 'ok!\n'

@app.route('/id/<id>')
def hornet(id):
    data = fetch(str(id))
    #result = compute(data)

    return str(data) + '\n'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=PORT, debug=True)