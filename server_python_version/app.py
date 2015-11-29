from flask import Flask
from fetch import fetch
from settings import PORT
from subprocess import call

app = Flask(__name__)

FILE_PATH = '/home/honeycomb/SparkTeam/part-00000'

@app.route('/result')
def result():
    return FILE_PATH

@app.route('/run')
def run():
    call(['/bin/spark-submit', '/home/honeycomb/SparkTeam/PySpark.py', '/user/honeycomb/sparkteam/input/sample_multiclass_classification_data.txt', '/user/honeycomb/sparkteam/input/sample_multiclass_classification_data_test.txt', '/home/honeycomb/SparkTeam'])
    with open(FILE_PATH) as fin:
        return fin.read()

@app.route('/id/<id>')
def hornet(id):
    data = fetch(str(id))
    #result = compute(data)

    return str(data) + '\n'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=PORT, debug=True)
