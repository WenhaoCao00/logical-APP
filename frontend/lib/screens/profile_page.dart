import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:quizeit/providers/user_state.dart';

class ProfilePage extends ConsumerWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final userState = ref.watch(userProvider);
    final userStateNotifier = ref.read(userProvider.notifier);

    void logout() {
      userStateNotifier.logout();
      context.go('/login');
    }

    void showEditDialog(BuildContext context, String title, String initialValue,
        Function(String) onSave) {
      final TextEditingController controller =
          TextEditingController(text: initialValue);

      showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text(title),
            content: TextField(
              controller: controller,
              decoration: InputDecoration(hintText: "Enter new $title"),
            ),
            actions: <Widget>[
              TextButton(
                child: const Text('Cancel'),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
              TextButton(
                child: const Text('Save'),
                onPressed: () {
                  onSave(controller.text);
                  Navigator.of(context).pop();
                },
              ),
            ],
          );
        },
      );
    }

    void showLogoutDialog(BuildContext context) {
      showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: const Text('Logout'),
            content: const Text('Are you sure you want to logout?'),
            actions: <Widget>[
              TextButton(
                child: const Text('Cancel'),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
              TextButton(
                child: const Text('Logout'),
                onPressed: () {
                  Navigator.of(context).pop();
                  logout();
                },
              ),
            ],
          );
        },
      );
    }

    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: const Text(
          'Profile Settings',
          style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
        ),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            GestureDetector(
              onLongPress: () {
                showLogoutDialog(context);
              },
              child: const CircleAvatar(
                radius: 40,
                child: Icon(
                  Icons.person,
                  size: 40,
                ),
              ),
            ),
            const SizedBox(height: 16),
            ListTile(
              title: const Text('Name'),
              subtitle: Text(userState.name),
              trailing: IconButton(
                icon: const Icon(Icons.edit),
                onPressed: () {
                  showEditDialog(
                    context,
                    'Name',
                    userState.name,
                    (newName) => userStateNotifier.updateUser(name: newName),
                  );
                },
              ),
            ),
            ListTile(
              title: const Text('Username'),
              subtitle: Text('@${userState.username}'),
            ),
            ListTile(
              title: const Text('Email'),
              subtitle: Text(userState.email),
            ),
            ListTile(
              title: const Text('Links'),
              subtitle: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  for (var url in userState.websiteUrls ?? [])
                    GestureDetector(
                      onLongPress: () {
                        showDialog(
                          context: context,
                          builder: (BuildContext context) {
                            return AlertDialog(
                              title: const Text('Delete Website'),
                              content: const Text(
                                  'Are you sure you want to delete this website?'),
                              actions: [
                                TextButton(
                                  onPressed: () {
                                    Navigator.of(context).pop();
                                  },
                                  child: const Text('Cancel'),
                                ),
                                TextButton(
                                  onPressed: () {
                                    userStateNotifier.removeWebsiteUrl(url);
                                    Navigator.of(context).pop();
                                  },
                                  child: const Text('Delete'),
                                ),
                              ],
                            );
                          },
                        );
                      },
                      child: Text(
                        url,
                        style: const TextStyle(
                          color: Colors.blue,
                          decoration: TextDecoration.underline,
                        ),
                      ),
                    ),
                  TextButton(
                    onPressed: () {
                      final TextEditingController linkController =
                          TextEditingController();
                      bool isValid = true;

                      showDialog(
                        context: context,
                        builder: (BuildContext context) {
                          return StatefulBuilder(
                            builder:
                                (BuildContext context, StateSetter setState) {
                              return AlertDialog(
                                title: const Text('Add Link'),
                                content: Column(
                                  mainAxisSize: MainAxisSize.min,
                                  children: [
                                    TextField(
                                      controller: linkController,
                                      decoration: InputDecoration(
                                        hintText: "Enter link here",
                                        errorText: isValid
                                            ? null
                                            : 'Link must start with https',
                                      ),
                                    ),
                                  ],
                                ),
                                actions: <Widget>[
                                  TextButton(
                                    onPressed: () {
                                      Navigator.of(context).pop();
                                    },
                                    child: const Text('Cancel'),
                                  ),
                                  TextButton(
                                    onPressed: () {
                                      if (linkController.text
                                          .startsWith('https')) {
                                        // Link is valid, handle the submitted link
                                        userStateNotifier
                                            .addWebsiteUrl(linkController.text);
                                        Navigator.of(context).pop();
                                      } else {
                                        // Link is not valid, show error
                                        setState(() {
                                          isValid = false;
                                        });
                                      }
                                    },
                                    child: const Text('Add'),
                                  ),
                                ],
                              );
                            },
                          );
                        },
                      );
                    },
                    child: const Text('+ Add link'),
                  ),
                ],
              ),
            ),
            ListTile(
              title: const Text('Bio'),
              subtitle: Text(userState.bio ?? ''),
              trailing: IconButton(
                icon: const Icon(Icons.edit),
                onPressed: () {
                  showEditDialog(
                    context,
                    'Bio',
                    userState.bio ?? '',
                    (newBio) => userStateNotifier.updateUser(bio: newBio),
                  );
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}
