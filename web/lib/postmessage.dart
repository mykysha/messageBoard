import 'dart:convert';

import 'package:http/http.dart';
import 'package:web/message.dart';

const uri = "http://localhost:8080/v1/messages";

Future<void> createServerMessage(
    String author, String text, String time) async {
  final client = Client();
  final request = await client.post(
    Uri.parse(uri),
    body: jsonEncode({
      'author': author,
      'text': text,
      'time': time,
    }),
  );
}

Future<List<Message>> getAllMessages() async {
  List<Message> messages = <Message>[];
  List<dynamic> data;
  final client = Client();
  final request = await client.get(
    Uri.parse(uri),
  );

  data = jsonDecode(request.body.toString());

  for (var i = 0; i < data.length; i++) {
    var author = data[i]["Author"] ?? "";
    var message = data[i]["Text"] ?? "";
    var time = data[i]["Time"] ?? "";

    messages.add(Message(author: author, message: message, time: time));
  }

  return messages;
}
