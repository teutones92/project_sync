import 'package:flutter/material.dart';
import 'package:project_sync/core/config/bloc_config.dart';

class Settings extends StatelessWidget {
  const Settings({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Padding(
      padding: const EdgeInsets.all(8.0),
      child: BlocBuilder<UserDataBloc, UserState?>(
        builder: (context, state) {
          final userBLoc = context.read<UserDataBloc>();
          return CustomScrollView(
            slivers: [
              SliverAppBar(
                title: const Text("Settings"),
                foregroundColor: Theme.of(context).brightness == Brightness.dark
                    ? Colors.white
                    : Colors.black,
                backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                floating: true,
                snap: true,
                centerTitle: true,
              ),
              state != null
                  ? SliverList(
                      delegate: SliverChildListDelegate(
                        [
                          Row(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              const CircleAvatar(radius: 80),
                              const SizedBox(width: 40),
                              Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                mainAxisAlignment: MainAxisAlignment.center,
                                children: [
                                  const Text("User Name"),
                                  const SizedBox(height: 8),
                                  ElevatedButton(
                                    style: ElevatedButton.styleFrom(
                                      elevation: 0,
                                      backgroundColor: Colors.transparent,
                                      side:
                                          const BorderSide(color: Colors.grey),
                                    ),
                                    onPressed: () {},
                                    child: const Text("Change Profile Picture"),
                                  ),
                                ],
                              ),
                            ],
                          ),
                          ListTile(
                            leading: const Icon(Icons.dark_mode),
                            title: const Text("Dark Mode"),
                            trailing: BlocBuilder<DarkModeBloc, ThemeMode>(
                              builder: (context, state) {
                                return Switch(
                                  value: state == ThemeMode.dark,
                                  onChanged: (value) {
                                    context
                                        .read<DarkModeBloc>()
                                        .setDarkMode(value);
                                    userBLoc.updateDarkModeUserData(
                                      value: value,
                                      context: context,
                                    );
                                  },
                                );
                              },
                            ),
                          ),
                          // ListTile(
                          //   title: const Text("Logout"),
                          //   onTap: () {},
                          // ),
                        ],
                      ),
                    )
                  : const SliverFillRemaining(
                      child: Center(
                        child: CircularProgressIndicator(),
                      ),
                    )
            ],
          );
        },
      ),
    ));
  }
}
