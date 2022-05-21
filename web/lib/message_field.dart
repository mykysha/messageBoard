import 'package:flutter/material.dart';

class MessageField extends StatelessWidget {
  const MessageField({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: MediaQuery.of(context).size.width,
      color: Colors.white60,
      alignment: Alignment.bottomCenter,
      child: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints constraints) {
          return Column(
            children: [
              Row(
                children: [
                  Container(
                    alignment: Alignment.centerLeft,
                    padding: const EdgeInsetsDirectional.fromSTEB(20, 20, 0, 0),
                    child: IntrinsicWidth(
                      child: TextField(
                        keyboardType: TextInputType.name,
                        maxLines: 1,
                        decoration: InputDecoration(
                          constraints: const BoxConstraints(minWidth: 94),
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(15),
                          ),
                          labelText: "nickname",
                        ),
                        onChanged: (text) {
                          print('Name text field: $text');
                        },
                      ),
                    ),
                  ),
                ],
              ),
              Container(
                padding: const EdgeInsets.all(20),
                width: constraints.maxWidth,
                child: TextField(
                  keyboardType: TextInputType.multiline,
                  maxLines: 5,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(15),
                    ),
                    labelText: "message",
                    suffixIcon: IconButton(
                      icon: const Icon(Icons.send_rounded),
                      onPressed: () {},
                    ),
                  ),
                  onChanged: (text) {
                    print('Message text field: $text');
                  },
                ),
              ),
            ],
          );
        },
      ),
    );
  }
}
