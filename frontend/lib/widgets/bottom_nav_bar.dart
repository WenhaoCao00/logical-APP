import 'package:flutter/material.dart';

class BottomNavBar extends StatelessWidget {
  final int currentIndex;
  final ValueChanged<int> onTap;

  const BottomNavBar(
      {super.key, required this.currentIndex, required this.onTap});

  @override
  Widget build(BuildContext context) {
    return BottomAppBar(
      shape: const CircularNotchedRectangle(),
      notchMargin: 6.0,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: <Widget>[
          IconButton(
            icon: Icon(
              Icons.person,
              color: currentIndex == 0 ? Colors.blue : Colors.grey,
            ),
            onPressed: () => onTap(0),
          ),
          IconButton(
            icon: Icon(
              Icons.list,
              color: currentIndex == 1 ? Colors.blue : Colors.grey,
            ),
            onPressed: () => onTap(1),
          ),
          IconButton(
            icon: Icon(
              Icons.quiz,
              color: currentIndex == 2 ? Colors.blue : Colors.grey,
            ),
            onPressed: () => onTap(2),
          ),
          IconButton(
            icon: Icon(
              Icons.question_answer,
              color: currentIndex == 3 ? Colors.blue : Colors.grey,
            ),
            onPressed: () => onTap(3),
          )
        ],
      ),
    );
  }
}
