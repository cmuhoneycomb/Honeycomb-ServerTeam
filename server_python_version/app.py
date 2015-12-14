from flask import Flask
from fetch import fetch
from settings import PORT
from subprocess import call
import time
import os

app = Flask(__name__)

FILE_PATH = '/home/honeycomb/SparkTeam/part-00000'

@app.route('/path')
def path():
    return FILE_PATH

@app.route('/result')
def result():
    return '[{"Order":"1","CLass":"Yes","Name":"TP Rate","Value":0.8095238095238095},{"Order":"2","CLass":"Yes","Name":"FP Rate","Value":0.10810810810810811},{"Order":"3","CLass":"Yes","Name":"Precision","Value":0.8095238095238095},{"Order":"4","CLass":"Yes","Name":"Recall","Value":0.8095238095238095},{"Order":"5","CLass":"Yes","Name":"F-Measure","Value":0.8095238095238095},{"Order":"1","CLass":"Yes","Name":"TP Rate","Value":1.0},{"Order":"2","CLass":"Yes","Name":"FP Rate","Value":0.0},{"Order":"3","CLass":"Yes","Name":"Precision","Value":1.0},{"Order":"4","CLass":"Yes","Name":"Recall","Value":1.0},{"Order":"5","CLass":"Yes","Name":"F-Measure","Value":1.0},{"Order":"1","CLass":"Yes","Name":"TP Rate","Value":0.7894736842105263},{"Order":"2","CLass":"Yes","Name":"FP Rate","Value":0.10256410256410256},{"Order":"3","CLass":"Yes","Name":"Precision","Value":0.7894736842105263},{"Order":"4","CLass":"Yes","Name":"Recall","Value":0.7894736842105263},{"Order":"5","CLass":"Yes","Name":"F-Measure","Value":0.7894736842105263}]'

@app.route('/run')
def run():
    call(['/bin/spark-submit', '/home/honeycomb/SparkTeam/PySpark.py', '/user/honeycomb/sparkteam/input/sample_multiclass_classification_data.txt', '/user/honeycomb/sparkteam/input/sample_multiclass_classification_data_test.txt', '/home/honeycomb/SparkTeam'])
    while not os.path.exists(FILE_PATH):
        time.sleep(1)
    with open(FILE_PATH) as fin:
        str =  fin.read()
    str = str.replace('\n', ',')
    return '[' +  str[:-1] + ']'

@app.route('/id/<id>')
def hornet(id):
    data = fetch(str(id))
    #result = compute(data)

    return str(data) + '\n'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=PORT, debug=True, threaded=True)