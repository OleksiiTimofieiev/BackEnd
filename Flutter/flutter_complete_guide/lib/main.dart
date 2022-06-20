import 'package:flutter/material.dart';

// void main() {
//   runApp(MyApp());
// }

void main() => runApp(MyApp());

// Can be rebuild
class MyApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    return MyAppState();
  }
}

// Persistent, data not being reset
class MyAppState extends State<MyApp> {
  var questionIndex = 0;

  void answerQuestion() {
    setState((){
      questionIndex = questionIndex + 1; 
    });
    print('Answer Chosen');
  }

  @override
  Widget build(BuildContext ctx) {
    var questions = [
      'Question 1',
      'Question 2',
    ];

    return MaterialApp(
      // home: Text('Hello !!!'),
      home: Scaffold(
        // basic page styling
        appBar: AppBar(
          title: Text('MyFirstApp'),
        ),
        body: Column(
          children: <Widget>[
            Text(questions.elementAt(questionIndex)),
            RaisedButton(child: Text('Answer 1'), onPressed: answerQuestion),
            RaisedButton(
                child: Text('Answer 2'),
                onPressed: () {
                  //
                  print('Answer 2');
                }),
            RaisedButton(
                child: Text('Answer 3'), onPressed: () => print("Answer 3")),
          ],
        ),
      ),
    );
  }
}
