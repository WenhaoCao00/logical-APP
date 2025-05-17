import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:quizeit/screens/signup_screen.dart';
import 'package:quizeit/screens/login_page.dart';
import 'package:quizeit/screens/quiz_page.dart';
import 'package:quizeit/screens/profile_page.dart';
import 'package:quizeit/screens/quiz_list_page.dart';
import 'package:quizeit/screens/admin_page.dart';
import 'package:quizeit/screens/guide_page.dart';
import 'package:quizeit/screens/quiz_result_page.dart';
import 'package:quizeit/widgets/scaffold_with_nested_navigation.dart';

final _rootNavigatorKey = GlobalKey<NavigatorState>();

final goRouter = GoRouter(
  navigatorKey: _rootNavigatorKey,
  initialLocation: '/login',
  routes: [
    // Login and Signup routes
    GoRoute(
      path: '/login',
      pageBuilder: (context, state) => const NoTransitionPage(
        child: LoginScreen(),
      ),
    ),
    GoRoute(
      path: '/signup',
      pageBuilder: (context, state) => const NoTransitionPage(
        child: SignupScreen(),
      ),
    ),
    // Main app routes with navigation bar
    StatefulShellRoute.indexedStack(
      builder: (context, state, navigationShell) {
        return ScaffoldWithNestedNavigation(navigationShell: navigationShell);
      },
      branches: [
        StatefulShellBranch(
          routes: [
            GoRoute(
              path: '/profile',
              pageBuilder: (context, state) => const NoTransitionPage(
                child: ProfilePage(),
              ),
            ),
          ],
        ),
        StatefulShellBranch(
          routes: [
            GoRoute(
              path: '/admin',
              pageBuilder: (context, state) => const NoTransitionPage(
                child: AdminPage(),
              ),
            ),
          ],
        ),
        StatefulShellBranch(
          routes: [
            GoRoute(
              path: '/quiz_list',
              pageBuilder: (context, state) => const NoTransitionPage(
                child: QuizListPage(),
              ),
            ),
            GoRoute(
              path: '/quiz_list/quiz/:quizId',
              builder: (context, state) {
                final quizId = int.parse(state.pathParameters['quizId']!);
                return QuizPage(quizId: quizId);
              },
            ),
            GoRoute(
              path: '/quiz_result',
              builder: (context, state) {
                final extra = state.extra as Map<String, dynamic>;
                final selectedOptionsString =
                    extra['selectedOptionsString'] as List<String>;
                final correctOptionsString =
                    extra['correctOptionsString'] as List<String>;
                final quizId = extra['quizId'] as int;

                return QuizResultPage(
                  selectedOptionsString: selectedOptionsString,
                  correctOptionsString: correctOptionsString,
                  quizId: quizId,
                );
              },
            ),
          ],
        ),
        StatefulShellBranch(
          routes: [
            GoRoute(
              path: '/guide',
              pageBuilder: (context, state) => const NoTransitionPage(
                child: GuidePage(),
              ),
            ),
          ],
        ),
      ],
    ),
  ],
);
