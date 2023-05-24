from flask import Flask, render_template, request
from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField


class LoginForm(FlaskForm):
    username = StringField('ユーザー名')
    password = PasswordField('パスワード')

app = Flask(__name__)
app.config['SECRET_KEY'] = 'secret'

@app.route('/', methods=['GET', 'POST'])
def index():
    form = LoginForm()
    if request.method == 'POST':
        return f'<h1>ユーザー名は {form.username.data}. パスワードは {form.password.data}.</h1>'
    return render_template('index.html', form=form)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=10000)