import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:quizeit/providers/quiz_state.dart';
import 'package:go_router/go_router.dart';
import 'dart:convert';

class QuizResultPage extends ConsumerStatefulWidget {
  final List<String> selectedOptionsString;
  final List<String> correctOptionsString;
  final int quizId;

  const QuizResultPage({
    required this.selectedOptionsString,
    required this.correctOptionsString,
    required this.quizId,
    super.key,
  });

  @override
  QuizResultPageState createState() => QuizResultPageState();
}

class QuizResultPageState extends ConsumerState<QuizResultPage> {
  @override
  Widget build(BuildContext context) {
    final quizState = ref.watch(quizProvider);

    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Quiz #${widget.quizId} Results',
              style: const TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),
            Expanded(
              child: ListView.builder(
                itemCount: quizState.questions.length,
                itemBuilder: (context, index) {
                  final question = quizState.questions[index];
                  final selectedOption = utf8
                      .decode(widget.selectedOptionsString[index].codeUnits);
                  final correctOption =
                      utf8.decode(widget.correctOptionsString[index].codeUnits);

                  return Card(
                    child: ListTile(
                      title: Text(
                        utf8.decode(question.question.runes.toList()),
                        style: const TextStyle(fontSize: 18),
                      ),
                      subtitle: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            'Your Answer: $selectedOption',
                            style: TextStyle(
                              color: selectedOption == correctOption
                                  ? Colors.green
                                  : Colors.red,
                            ),
                          ),
                          Text(
                            'Correct Answer: $correctOption',
                            style: const TextStyle(color: Colors.green),
                          ),
                        ],
                      ),
                    ),
                  );
                },
              ),
            ),
            const SizedBox(height: 16),
            Center(
              child: ElevatedButton(
                onPressed: () {
                  context.go('/quiz_list');
                },
                child: const Text('Back to Quiz List'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
