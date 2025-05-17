import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

class QuizState {
  final List<Question> questions;
  QuizState({required this.questions});
}

class QuizNotifier extends StateNotifier<QuizState> {
  QuizNotifier() : super(QuizState(questions: []));

  Future<void> fetchQuiz(String endpoint) async {
    final url = Uri.parse(endpoint);
    try {
      final response = await http.get(url);
      if (response.statusCode == 200) {
        List<dynamic> data = json.decode(response.body);
        List<Question> newQuestions =
            data.map((item) => Question.fromJson(item)).toList();
        state = QuizState(questions: newQuestions);
      } else {
        throw Exception('Failed to load quiz');
      }
    } catch (e) {
      throw Exception('Failed to load quiz: $e');
    }
  }

  Future<void> generateConditionalQuiz() async {
    await fetchQuiz('http://10.0.2.2:7999/generateConditionalTask/');
  }

  Future<void> generateDeMorganQuiz() async {
    await fetchQuiz('http://10.0.2.2:7999/generateDeMorganTask/');
  }

  Future<void> generateTruthTableQuiz() async {
    await fetchQuiz('http://10.0.2.2:7999/generateTruthTableTask/');
  }

  Future<void> generateEasyQuiz() async {
    await fetchQuiz('http://10.0.2.2:7999/getTask/easy');
  }

  Future<void> generateMediumQuiz() async {
    await fetchQuiz('http://10.0.2.2:7999/getTask/medium');
  }

  Future<void> generateHardQuiz() async {
    await fetchQuiz('http://10.0.2.2:7999/getTask/hard');
  }
}

final quizProvider = StateNotifierProvider<QuizNotifier, QuizState>((ref) {
  return QuizNotifier();
});

class Question {
  final String question;
  final List<String> options;
  final String correctAnswer;

  Question({
    required this.question,
    required this.options,
    required this.correctAnswer,
  });

  factory Question.fromJson(Map<String, dynamic> json) {
    List<String> options;

    // Check if the response contains 'wrong_answers' (old API format)
    if (json.containsKey('wrong_answers')) {
      options = [
        ...json['wrong_answers'].map((answer) => answer.toString()),
        json['correct_answer'].toString()
      ];
    } else {
      // New API format with full options list
      options = List<String>.from(json['options']);
    }

    options.shuffle(); // Shuffle to randomize the options

    return Question(
      question: json['question'].toString(),
      options: options,
      correctAnswer: json['correct_answer'].toString(),
    );
  }
}
