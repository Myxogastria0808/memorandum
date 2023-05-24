class Person:
    #初期値のオブジェクト
    ##self・・・呼び出したインスタンス自身
    def __init__(self, name, nationality, age):
        self.name = name
        self.nationality = nationality
        self.age = age
    def __call__(self):
        print('ここはcall関数です。')
    
    def say_hello(self, name):
        print("こんにちわ{}さん、私は{}です。".format(name, self.name))

#インスタンス化
imanishi = Person('今西', '日本', 25)
mike = Person('マイク', 'アメリカ', 13)

#属性の表示
print(imanishi.name)
print(imanishi.nationality)
print(imanishi.age)

#メソッドの実行
imanishi.say_hello('Alexa')
mike.say_hello('Alexa')

#call関数の呼び出し
#インスタンス名()で呼び出せる
imanishi()