import 'package:flutter/material.dart';

var exampleMessages = <Message>[
  const Message(
    author: "Winnie the Pooh",
    message: "I am short, fat, and proud of that!",
    time: "01/01/2021 12:59",
  ),
  const Message(
    author: "Piglet",
    message: "I think..."
        " I have just remembered something"
        " that I forgot to do yesterday"
        " and shan’t is able to do tomorrow...",
    time: "02/02/2020 13:59",
  ),
  const Message(
    author: "Tigger",
    message: "Climbing trees? Why, that’s what Tiggers do best!"
        " Only Tiggers don’t climb trees; they bounce ’em!",
    time: "03/03/2019 19:59",
  ),
];

class Message extends StatelessWidget {
  final String author;
  final String message;
  final String time;

  const Message(
      {Key? key,
      required this.author,
      required this.message,
      required this.time})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: MediaQuery.of(context).size.width,
      padding: const EdgeInsetsDirectional.fromSTEB(20, 20, 20, 0),
      child: Container(
        padding: const EdgeInsetsDirectional.all(10),
        decoration: const BoxDecoration(
          color: Colors.blueAccent,
          borderRadius: BorderRadius.all(Radius.circular(15)),
        ),
        child: Column(
          children: <Widget>[
            SizedBox(
              child: Row(
                children: <Widget>[
                  Expanded(
                    child: Align(
                      alignment: Alignment.topLeft,
                      child: Text(
                        author != "" ? author : "anonymous",
                        softWrap: true,
                        style: const TextStyle(
                          color: Colors.orange,
                        ),
                      ),
                    ),
                  ),
                  Align(
                    alignment: Alignment.topRight,
                    child: Text(
                      time,
                      softWrap: true,
                      style: const TextStyle(
                        color: Colors.orange,
                      ),
                    ),
                  ),
                ],
              ),
            ),
            Container(
              height: 10,
            ),
            Align(
              alignment: Alignment.topLeft,
              child: Text(
                message,
                softWrap: true,
                style: const TextStyle(
                  color: Colors.black,
                  fontSize: 20,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
