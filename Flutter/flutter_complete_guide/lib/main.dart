import 'package:flutter/material.dart';
import 'package:flutter_complete_guide/answer.dart';
import './question.dart';

// void main() {
//   runApp(MyApp());
// }

void main() => runApp(MyApp());

// Can be rebuild
class MyApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    return _MyAppState();
  }
}

// Persistent, data not being reset
class _MyAppState extends State<MyApp> {
  var _questionIndex = 0;

  void _answerQuestion() {
    setState(() {
      _questionIndex = _questionIndex + 1;
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
            Question(questions.elementAt(_questionIndex)),
            Answer(_answerQuestion),
            Answer(_answerQuestion),
            Answer(_answerQuestion),
          ],
        ),
      ),
    );
  }
}
