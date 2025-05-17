import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class UserState {
  final String name;
  final String username;
  final String email;
  final String? avatarUrl;
  final List<String>? websiteUrls;
  final String? bio;
  final String? description;
  final int? totalQuizNumber;
  final Map<String, int>? quizScores;

  UserState({
    required this.name,
    required this.username,
    required this.email,
    this.avatarUrl,
    this.websiteUrls,
    this.bio,
    this.description,
    this.totalQuizNumber,
    this.quizScores,
  });

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'username': username,
      'email': email,
      'avatarUrl': avatarUrl,
      'websiteUrls': websiteUrls,
      'bio': bio,
      'description': description,
      'totalQuizNumber': totalQuizNumber,
      'quizScores': quizScores,
    };
  }

  factory UserState.fromJson(Map<String, dynamic> json) {
    return UserState(
      name: json['name'],
      username: json['username'],
      email: json['email'],
      avatarUrl: json['avatarUrl'],
      websiteUrls: json['websiteUrls'] != null
          ? List<String>.from(json['websiteUrls'])
          : null,
      bio: json['bio'],
      description: json['description'],
      totalQuizNumber: json['totalQuizNumber'],
      quizScores: json['quizScores'] != null
          ? Map<String, int>.from(json['quizScores'])
          : null,
    );
  }

  UserState copyWith({
    String? name,
    String? username,
    String? email,
    String? avatarUrl,
    List<String>? websiteUrls,
    String? bio,
    String? description,
    int? totalQuizNumber,
    Map<String, int>? quizScores,
  }) {
    return UserState(
      name: name ?? this.name,
      username: username ?? this.username,
      email: email ?? this.email,
      avatarUrl: avatarUrl ?? this.avatarUrl,
      websiteUrls: websiteUrls ?? this.websiteUrls,
      bio: bio ?? this.bio,
      description: description ?? this.description,
      totalQuizNumber: totalQuizNumber ?? this.totalQuizNumber,
      quizScores: quizScores ?? this.quizScores,
    );
  }
}

class UserNotifier extends StateNotifier<UserState> {
  String? _password;

  UserNotifier()
      : super(UserState(
          name: '',
          username: '',
          email: '',
          avatarUrl: null,
          websiteUrls: [],
          bio: null,
          description: null,
          totalQuizNumber: 0,
          quizScores: {},
        ));

  Future<void> _loadUserFromBackend() async {
    if (state.username.isEmpty || _password == null) return;

    final url = Uri.parse('http://10.0.2.2:7999/getUserState/');
    final response = await http.get(
      url,
      headers: {
        'Authorization':
            'Basic ${base64Encode(utf8.encode('${state.username}:$_password'))}',
        'Content-Type': 'application/json',
      },
    );

    if (response.statusCode == 200) {
      final json = jsonDecode(response.body);
      state = UserState.fromJson(json);
    } else {
      throw Exception('Failed to load user state from backend');
    }
  }

  Future<void> signup(
      String name, String username, String email, String password) async {
    final url = Uri.parse('http://10.0.2.2:7999/signup/');
    final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'name': name,
        'username': username,
        'email': email,
        'password': password,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Signup failed');
    }
  }

  Future<void> login(String email, String password) async {
    try {
      final response = await http.post(
        Uri.parse('http://10.0.2.2:7999/login/'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'email': email, 'password': password}),
      );

      if (response.statusCode == 200) {
        final json = jsonDecode(response.body);
        _password = password;
        state = UserState(
          name: json['name'],
          username: json['username'],
          email: json['email'],
          avatarUrl: null,
          websiteUrls: json['websiteUrls'] != null
              ? List<String>.from(json['websiteUrls'])
              : [],
          bio: null,
          description: null,
          totalQuizNumber: 0,
          quizScores: {},
        );

        await _loadUserFromBackend();

        //print(all states)
      } else {
        throw Exception('Failed to login');
      }
    } catch (e) {
      throw Exception('Failed to login');
    }
  }

  void logout() {
    _password = null;
    state = UserState(
      name: '',
      username: '',
      email: '',
      avatarUrl: null,
      websiteUrls: [],
      bio: null,
      description: null,
      totalQuizNumber: 0,
      quizScores: {},
    );
  }

  void updateUser({
    String? name,
    String? username,
    String? email,
    String? avatarUrl,
    List<String>? websiteUrls,
    String? bio,
    String? description,
    int? totalQuizNumber,
    Map<String, int>? quizScores,
  }) {
    state = state.copyWith(
      name: name,
      username: username,
      email: email,
      avatarUrl: avatarUrl,
      websiteUrls: websiteUrls,
      bio: bio,
      description: description,
      totalQuizNumber: totalQuizNumber,
      quizScores: quizScores,
    );
    _syncWithBackend();
  }

  int get totalScore {
    if (state.quizScores == null) return 0;
    return state.quizScores!.values.fold(0, (a, b) => a + b);
  }

  int get highestScore {
    if (state.quizScores == null || state.quizScores!.isEmpty) return 0;
    return state.quizScores!.values.fold(0, (a, b) => a > b ? a : b);
  }

  void incrementQuizNumber() {
    state = state.copyWith(totalQuizNumber: (state.totalQuizNumber ?? 0) + 1);
    _syncWithBackend(); // changed
  }

  void setQuizScore(String quizId, int score) {
    final updatedQuizScores = Map<String, int>.from(state.quizScores ?? {});
    updatedQuizScores[quizId] = score;
    state = state.copyWith(quizScores: updatedQuizScores);
    _syncWithBackend();
  }

  void removeWebsiteUrl(String url) {
    final updatedWebsiteUrls = List<String>.from(state.websiteUrls ?? []);
    updatedWebsiteUrls.remove(url);
    state = state.copyWith(websiteUrls: updatedWebsiteUrls);
    _syncWithBackend();
  }

  void addWebsiteUrl(String url) {
    final updatedWebsiteUrls = List<String>.from(state.websiteUrls ?? []);
    updatedWebsiteUrls.add(url);
    state = state.copyWith(websiteUrls: updatedWebsiteUrls);
    _syncWithBackend();
  }

  Future<void> _syncWithBackend() async {
    if (state.username.isEmpty || _password == null) return;

    final url = Uri.parse('http://10.0.2.2:7999/updateUserState/');
    final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization':
            'Basic ${base64Encode(utf8.encode('${state.username}:$_password'))}',
      },
      body: jsonEncode(state.toJson()),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to sync with backend');
    }
  }
}

final userProvider = StateNotifierProvider<UserNotifier, UserState>((ref) {
  return UserNotifier();
});
