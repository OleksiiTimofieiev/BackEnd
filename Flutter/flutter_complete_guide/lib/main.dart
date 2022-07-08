import 'package:flutter/material.dart';
import 'package:flutter_complete_guide/answer.dart';
import './quiz.dart';
import './result.dart';

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
  var _totalScore = 0;

  void _resetQuiz() {
    setState(() {
      _questionIndex = 0;
      _totalScore = 0;
    });
  }

  final _questions = [
    {
      'questionText': 'Fav color ?',
      'answers': [
        {'text': 'Black', 'score': 10},
        {'text': 'Red', 'score': 6},
        {'text': 'Green', 'score': 3},
        {'text': 'White', 'score': 1},
      ],
    },
    {
      'questionText': 'Fav animal ?',
      'answers': [
        {'text': 'Rabbit', 'score': 1},
        {'text': 'Snacke', 'score': 1},
        {'text': 'Elefant', 'score': 1},
        {'text': 'Lion', 'score': 1},
      ],
    },
    {
      'questionText': 'Fav instructor ?',
      'answers': [
        {'text': 'Max', 'score': 1},
        {'text': 'Maxx', 'score': 1},
        {'text': 'Max', 'score': 1},
        {'text': 'Max', 'score': 1},
      ],
    }
  ];

  void _answerQuestion(int score) {
    _totalScore += score;

    setState(() {
      if (_questionIndex < _questions.length) {
        _questionIndex = _questionIndex + 1;
      }
    });
    print('Answer Chosen');
  }

  @override
  Widget build(BuildContext ctx) {
    return MaterialApp(
      // home: Text('Hello !!!'),
      home: Scaffold(
        // basic page styling
        appBar: AppBar(
          title: Text('MyFirstApp'),
        ),
        body: _questionIndex < _questions.length
            ? Quiz(
                answerQuestion: _answerQuestion,
                questionIndex: _questionIndex,
                questions: _questions,
              )
            : Result(_totalScore, _resetQuiz),
      ),
    );
  }
}
