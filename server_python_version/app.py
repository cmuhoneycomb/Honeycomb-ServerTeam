
from flask import Flask

app = Flask(__name__)

@app.route('/test')
def test():
    return 'ok!\n'

@app.route('/id/<id>')
def hornet(id):
    data = fetch(str(id))
    #result = compute(data)
    
    return str(data) + '\n'
    return data

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=32769, debug=True)
