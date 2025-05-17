import 'package:flutter/material.dart';

class GuidePage extends StatelessWidget {
  const GuidePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Welcome to the App Guide'),
      ),
      body: const SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                'This guide is designed to help you navigate through the main features of our app. Each section below highlights a key screen with a brief description and a corresponding screenshot.',
                style: TextStyle(fontSize: 16),
              ),
              SizedBox(height: 32),

              // Quiz List Section
              Text(
                '1. Quiz List',
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 8),
              Text(
                'The Quiz List screen displays all the quizzes available to you. For each quiz, you can see its title and status, such as "Not started" or your current score. '
                'To start a quiz, simply click on the "Start" button next to the quiz name. This will take you directly to the quiz interface, where you can begin answering questions.',
                style: TextStyle(fontSize: 16),
              ),
              SizedBox(height: 8),
              Image(
                image: AssetImage('assets/pic/admin.jpg'), //
                fit: BoxFit.contain,
              ),
              SizedBox(height: 32),

              // Profile Settings Section
              Text(
                '2. Profile Settings',
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 8),
              Text(
                'The Profile Settings screen allows you to customize your personal information. Here, you can update your name, username, email, and bio. '
                'You can also add links to your social media profiles or personal website. To edit any of these details, click on the pencil icon next to the field you wish to change.',
                style: TextStyle(fontSize: 16),
              ),
              SizedBox(height: 8),
              Image(
                image: AssetImage('assets/pic/profile.jpg'),
                fit: BoxFit.contain,
              ),
              SizedBox(height: 32),

              // Admin Dashboard Section
              Text(
                '3. Statistic Dashboard',
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 8),
              Text(
                'The Statistic Dashboard is your go-to screen for viewing key statistics related to your quizzes. You can track the total number of quizzes you have taken, your correct answer rate, '
                'your highest score, and your cumulative score across all quizzes. Additionally, this screen provides reminders or important notices relevant to your quiz activity.',
                style: TextStyle(fontSize: 16),
              ),
              SizedBox(height: 8),
              Image(
                image: AssetImage('assets/pic/admin_dashboard.png'),
                fit: BoxFit.contain,
              ),
              SizedBox(height: 32),

              // Quiz Page Section
              Text(
                '4. Quiz Page',
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 8),
              Text(
                'The Quiz Page is where you actually take the quizzes. It displays the question at the top, followed by multiple choice answers. '
                'You can select an answer and navigate between questions using the "Previous" and "Next" buttons. Make sure to answer all questions before the timer runs out!',
                style: TextStyle(fontSize: 16),
              ),
              SizedBox(height: 8),
              Image(
                image: AssetImage('assets/pic/quiz_page.png'),
                fit: BoxFit.contain,
              ),
              SizedBox(height: 32),

              // Quiz Result Section
              Text(
                '5. Quiz Results',
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 8),
              Text(
                'After completing a quiz, you will be directed to the Quiz Results page. Here, you can see your selected answers compared to the correct answers. '
                'This page provides feedback on your performance, showing you what you got right and what you missed.',
                style: TextStyle(fontSize: 16),
              ),
              SizedBox(height: 8),
              Image(
                image: AssetImage('assets/pic/result.png'),
                fit: BoxFit.contain,
              ),
              SizedBox(height: 32),

              // Conclusion
              Text(
                'We hope this guide helps you navigate the app more efficiently. If you have any further questions, feel free to contact support.',
                style: TextStyle(fontSize: 16),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
