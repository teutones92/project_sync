import 'package:project_sync/core/config/bloc_config.dart';
import 'package:flutter/material.dart';

class UserAvatarAndName extends StatelessWidget {
  const UserAvatarAndName({super.key});

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: context.read<UserDataBloc>().getUserData(context),
        builder: (_, __) {
          return BlocBuilder<UserDataBloc, UserState?>(
            builder: (context, state) {
              context.read<DarkModeBloc>().setDarkMode(state?.darkMode ?? true);
              return AnimatedSwitcher(
                duration: const Duration(milliseconds: 500),
                child: state != null
                    ? Row(
                        crossAxisAlignment: CrossAxisAlignment.center,
                        children: [
                          InkWell(
                            borderRadius: BorderRadius.circular(10),
                            onTap: () {
                              // context.read<HomeBloc>().logOut(context);
                            },
                            child: Row(
                              children: [
                                const Icon(Icons.keyboard_arrow_down),
                                Text(
                                  state.username,
                                  style: const TextStyle(
                                      fontWeight: FontWeight.bold),
                                ),
                              ],
                            ),
                          ),
                          const SizedBox(width: 10),
                          CircleAvatar(
                            backgroundImage: state.userAvatarPath.isEmpty
                                ? null
                                : NetworkImage(
                                    state.userAvatarPath,
                                  ),
                            radius: 25,
                            child: state.userAvatarPath.isEmpty
                                ? Text(
                                    state.username[0],
                                    style: const TextStyle(
                                      fontWeight: FontWeight.bold,
                                      shadows: [
                                        Shadow(
                                          color: Colors.black,
                                          blurRadius: 2,
                                        ),
                                      ],
                                    ),
                                  )
                                : null,
                          ),
                        ],
                      )
                    : SizedBox(
                        width: 200,
                        height: 4,
                        child: LinearProgressIndicator(
                          backgroundColor: Colors.grey.withOpacity(0.5),
                          color: Colors.white.withOpacity(0.5),
                        ),
                      ),
              );
            },
          );
        });
  }
}
