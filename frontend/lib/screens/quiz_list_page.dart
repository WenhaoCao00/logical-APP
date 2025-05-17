import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:quizeit/providers/user_state.dart';
import 'package:quizeit/providers/quiz_state.dart';

class QuizListPage extends ConsumerWidget {
  const QuizListPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final userState = ref.watch(userProvider);
    return Scaffold(
      appBar: AppBar(
        title: const Text(
          'Quiz List',
          style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
        ),
      ),
      body: ListView(
        padding: const EdgeInsets.all(16.0),
        children: [
          const Text('Random Generated quiz',
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold)),
          _buildQuizTile(context, ref, 1, 'Conditional Quiz', userState),
          _buildQuizTile(context, ref, 2, 'De Morgan\'s Quiz', userState),
          _buildQuizTile(context, ref, 3, 'Truth Table Quiz', userState),
          const Text('Fixed quiz',
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold)),
          _buildQuizTile(context, ref, 4, 'Easy Quiz', userState),
          _buildQuizTile(context, ref, 5, 'Medium Quiz', userState),
          _buildQuizTile(context, ref, 6, 'Hard Quiz', userState),
        ],
      ),
    );
  }

  Widget _buildQuizTile(BuildContext context, WidgetRef ref, int quizNumber,
      String title, UserState userState) {
    final quizScores = userState.quizScores ?? {};

    return ListTile(
      title: Text(title),
      subtitle: Text(
        quizScores['quiz$quizNumber'] != null
            ? 'Score: ${quizScores['quiz$quizNumber']}'
            : 'Not started',
      ),
      trailing: ElevatedButton(
        onPressed: () {
          _showStartQuizDialog(context, ref, quizNumber, title);
        },
        child: const Text('Start'),
      ),
    );
  }

  void _showStartQuizDialog(
      BuildContext context, WidgetRef ref, int quizNumber, String title) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text('Start $title'),
          content: Text('Are you sure you want to start $title?'),
          actions: <Widget>[
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: const Text('Cancel'),
            ),
            TextButton(
              onPressed: () {
                Navigator.pop(context);
                _startQuiz(ref, quizNumber);
                context.push('/quiz_list/quiz/$quizNumber');
              },
              child: const Text('Start'),
            ),
          ],
        );
      },
    );
  }

  void _startQuiz(WidgetRef ref, int quizNumber) {
    final quizNotifier = ref.read(quizProvider.notifier);

    switch (quizNumber) {
      case 1:
        quizNotifier.generateConditionalQuiz();
        break;
      case 2:
        quizNotifier.generateDeMorganQuiz();
        break;
      case 3:
        quizNotifier.generateTruthTableQuiz();
        break;
      case 4:
        quizNotifier.generateEasyQuiz();
        break;
      case 5:
        quizNotifier.generateMediumQuiz();
        break;
      case 6:
        quizNotifier.generateHardQuiz();
        break;
      default:
        throw Exception('Invalid quiz number');
    }
  }
}
