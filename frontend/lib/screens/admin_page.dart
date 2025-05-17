import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import "package:quizeit/providers/user_state.dart";

class AdminPage extends ConsumerWidget {
  const AdminPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final userState = ref.watch(userProvider);
    final userStateNotifier = ref.read(userProvider.notifier);
    final int totalQuiz = userState.totalQuizNumber ?? 0;
    final int totalScore = userStateNotifier.totalScore;

    final int highestScore = userStateNotifier.highestScore;

    return Scaffold(
      appBar: AppBar(
          automaticallyImplyLeading: false,
          title: const Text(
            'Statistics',
            style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
          )),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            const SizedBox(height: 32),
            Card(
              child: ListTile(
                title: const Text('Total Finished Quizzes'),
                trailing: Text(totalQuiz.toString(),
                    style: const TextStyle(
                        fontSize: 18, fontWeight: FontWeight.bold)),
              ),
            ),
            Card(
              child: ListTile(
                title: const Text('Unfinished Quizzes'),
                trailing: Text((6 - totalQuiz).toString(),
                    style: const TextStyle(
                        fontSize: 18, fontWeight: FontWeight.bold)),
              ),
            ),
            Card(
              child: ListTile(
                title: const Text('Highest Score'),
                trailing: Text(highestScore.toString(),
                    style: const TextStyle(
                        fontSize: 18, fontWeight: FontWeight.bold)),
              ),
            ),
            Card(
              child: ListTile(
                title: const Text('Total Score'),
                trailing: Text(totalScore.toString(),
                    style: const TextStyle(
                        fontSize: 18, fontWeight: FontWeight.bold)),
              ),
            ),
            const SizedBox(height: 32),
            const Text(
              'Reminder',
              style: TextStyle(fontSize: 16, fontWeight: FontWeight.w500),
            ),
            Text(
              'It is time to start quiz!',
              style: TextStyle(fontSize: 14, color: Colors.grey[700]),
            ),
            Text(
              'admin@admin.com',
              style: TextStyle(fontSize: 14, color: Colors.grey[700]),
            ),
          ],
        ),
      ),
    );
  }
}
