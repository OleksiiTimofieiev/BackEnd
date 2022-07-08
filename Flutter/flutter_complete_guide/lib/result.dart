import 'package:flutter/src/foundation/key.dart';
import 'package:flutter/src/widgets/framework.dart';

import 'package:flutter/material.dart';

class Result extends StatelessWidget {
  final int resultScore;
  final Function resetFunction;

  Result(this.resultScore, this.resetFunction);

  String get resultPhrase {
    var resulText = 'WooHoo';

    if (resultScore <= 8) {
      resulText = 'Awesome';
    }
    return resulText;
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          Text(
            resultPhrase,
            style: TextStyle(fontSize: 36, fontWeight: FontWeight.bold),
            textAlign: TextAlign.center,
          ),
          FlatButton(onPressed: resetFunction, child: Text("RestartQuiz"), textColor: Colors.blue,)
        ],
      ),
    );
  }
}
