import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'dart:async';
import 'package:quizeit/providers/quiz_state.dart';
import 'package:quizeit/providers/user_state.dart';
import 'package:go_router/go_router.dart';
import 'dart:convert';

class QuizPage extends ConsumerStatefulWidget {
  final int quizId;
  const QuizPage({
    required this.quizId,
    super.key,
  });

  @override
  QuizPageState createState() => QuizPageState();
}

class QuizPageState extends ConsumerState<QuizPage> {
  int _selectedIndex = -1;
  int _currentQuestionIndex = 0;
  int _remainingTime = 600;

  List<int> _selectedOptions = [];

  late Timer _timer;

  @override
  void initState() {
    super.initState();

    _startTimer();

    WidgetsBinding.instance.addPostFrameCallback((_) async {
      final quizNotifier = ref.read(quizProvider.notifier);
      await Future<void>.delayed(Duration.zero);
      switch (widget.quizId) {
        case 1:
          await quizNotifier.generateConditionalQuiz();
          break;
        case 2:
          await quizNotifier.generateDeMorganQuiz();
          break;
        case 3:
          await quizNotifier.generateTruthTableQuiz();
          break;
        case 4:
          await quizNotifier.generateEasyQuiz();
          break;
        case 5:
          await quizNotifier.generateMediumQuiz();
          break;
        case 6:
          await quizNotifier.generateHardQuiz();
          break;
        default:
          await quizNotifier.generateConditionalQuiz();
      }

      final quizState = ref.read(quizProvider);
      setState(() {
        _selectedOptions = List<int>.filled(quizState.questions.length, -1);
      });
    });
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  void _startTimer() {
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      setState(() {
        if (_remainingTime > 0) {
          _remainingTime--;
        } else {
          _timer.cancel();
          _onSubmit();
        }
      });
    });
  }

  void _onOptionSelected(int index) {
    setState(() {
      _selectedIndex = index;
      _selectedOptions[_currentQuestionIndex] = index;
    });
  }

  void _onNextQuestion() {
    setState(() {
      if (_currentQuestionIndex < _selectedOptions.length - 1) {
        _currentQuestionIndex++;
        _selectedIndex = _selectedOptions[_currentQuestionIndex];
      }
    });
  }

  void _onPreviousQuestion() {
    setState(() {
      if (_currentQuestionIndex > 0) {
        _currentQuestionIndex--;
        _selectedIndex = _selectedOptions[_currentQuestionIndex];
      }
    });
  }

  void _onSubmit() {
    // 获取当前的题目列表
    final quizState = ref.read(quizProvider);
    final userNotifier = ref.read(userProvider.notifier);
    final userState = ref.read(userProvider);

    // 计算得分
    int score = 0;
    List<String> correctOptionsString = [];
    List<String> selectedOptionsString = [];

    for (int i = 0; i < quizState.questions.length; i++) {
      correctOptionsString.add(quizState.questions[i].correctAnswer);

      if (_selectedOptions[i] != -1) {
        selectedOptionsString
            .add(quizState.questions[i].options[_selectedOptions[i]]);
      } else {
        selectedOptionsString.add('No Selection');
      }

      if (_selectedOptions[i] != -1 &&
          quizState.questions[i].options[_selectedOptions[i]] ==
              quizState.questions[i].correctAnswer) {
        score++;
      }
    }
    userNotifier.setQuizScore('quiz${widget.quizId}', score);

    if ((userState.quizScores ?? {})['quiz${widget.quizId}'] == null) {
      userNotifier.incrementQuizNumber();
    }

    context.go(
      '/quiz_result',
      extra: {
        'selectedOptionsString': selectedOptionsString,
        'correctOptionsString': correctOptionsString,
        'quizId': widget.quizId,
      },
    );
  }

  String _formatTime(int seconds) {
    int minutes = seconds ~/ 60;
    int secs = seconds % 60;
    return '${minutes.toString().padLeft(2, '0')}:${secs.toString().padLeft(2, '0')}';
  }

  @override
  Widget build(BuildContext context) {
    final quizState = ref.watch(quizProvider);

    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Text(
          'Question #${_currentQuestionIndex + 1}',
          style: const TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
        ),
      ),
      body: SafeArea(
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                Text(
                  'Left Time: ${_formatTime(_remainingTime)}',
                  style: const TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                const SizedBox(height: 16),
                if (quizState.questions.isNotEmpty)
                  Card(
                    child: Container(
                      height: 150,
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(16.0),
                      child: Text(
                        utf8.decode(quizState
                            .questions[_currentQuestionIndex].question
                            .toString()
                            .runes
                            .toList()),
                        style: const TextStyle(fontSize: 18),
                      ),
                    ),
                  ),
                const SizedBox(height: 16),
                const Text(
                  'Please select one answer',
                  style: TextStyle(
                    fontSize: 14,
                    color: Colors.grey,
                  ),
                ),
                const SizedBox(height: 16),
                if (quizState.questions.isNotEmpty)
                  SizedBox(
                    height: 120,
                    child: GridView.builder(
                      gridDelegate:
                          const SliverGridDelegateWithFixedCrossAxisCount(
                        crossAxisCount: 2,
                        crossAxisSpacing: 10,
                        mainAxisSpacing: 20,
                        childAspectRatio: 4,
                      ),
                      itemCount: quizState
                          .questions[_currentQuestionIndex].options.length,
                      itemBuilder: (context, index) {
                        return ElevatedButton(
                          onPressed: () => _onOptionSelected(index),
                          style: ElevatedButton.styleFrom(
                            backgroundColor: _selectedIndex == index
                                ? Colors.black
                                : Colors.grey,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(20),
                            ),
                          ),
                          child: Text(
                            utf8.decode(quizState
                                .questions[_currentQuestionIndex].options[index]
                                .toString()
                                .runes
                                .toList()),
                            style: const TextStyle(
                                color: Colors.white, fontSize: 12),
                          ),
                        );
                      },
                    ),
                  ),
                const SizedBox(height: 16),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    if (_currentQuestionIndex > 0)
                      ElevatedButton.icon(
                        onPressed: _onPreviousQuestion,
                        icon: const Icon(Icons.arrow_back),
                        label: const Text('Previous'),
                      ),
                    const Spacer(),
                    if (_currentQuestionIndex < _selectedOptions.length - 1)
                      ElevatedButton.icon(
                        onPressed: _onNextQuestion,
                        icon: const Icon(Icons.arrow_forward),
                        label: const Text('Next'),
                      ),
                    if (_currentQuestionIndex == _selectedOptions.length - 1)
                      ElevatedButton(
                        onPressed: () {
                          showDialog(
                            context: context,
                            builder: (BuildContext context) {
                              return AlertDialog(
                                title: const Text('Confirm Submission'),
                                content: const Text(
                                    'Are you sure you want to submit?'),
                                actions: <Widget>[
                                  TextButton(
                                    onPressed: () {
                                      Navigator.of(context)
                                          .pop(); // Dismiss the dialog
                                    },
                                    child: const Text('Cancel'),
                                  ),
                                  ElevatedButton(
                                    onPressed: () {
                                      Navigator.of(context)
                                          .pop(); // Dismiss the dialog
                                      _onSubmit(); // Call the submit function
                                    },
                                    child: const Text('Submit'),
                                  ),
                                ],
                              );
                            },
                          );
                        },
                        child: const Text('Submit'),
                      ),
                  ],
                ),
                const SizedBox(height: 16),
                const Text(
                  'Question Status:',
                  style: TextStyle(fontSize: 18, fontWeight: FontWeight.w500),
                ),
                const SizedBox(height: 16),
                SizedBox(
                  height: 100,
                  child: GridView.count(
                    crossAxisCount: 10,
                    crossAxisSpacing: 8.0,
                    childAspectRatio: 3,
                    children: List.generate(_selectedOptions.length, (index) {
                      return Container(
                        padding: const EdgeInsets.all(8.0),
                        decoration: BoxDecoration(
                          color: _selectedOptions[index] == -1
                              ? Colors.grey
                              : Colors.black,
                          borderRadius: BorderRadius.circular(8),
                        ),
                        child: Center(
                          child: Text(
                            _selectedOptions[index] == -1
                                ? 'Not Selected'
                                : 'Selected',
                            style: const TextStyle(color: Colors.white),
                            textAlign: TextAlign.center,
                          ),
                        ),
                      );
                    }),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
