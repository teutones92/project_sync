import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/views/0.1_home/widgets/home_header/widgets/user_avatar_and_name.dart';
import 'package:flutter/material.dart';

import 'widgets/notification_button.dart';

class HomeHeader extends StatelessWidget {
  const HomeHeader({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<AnimateSideMenuBloc, double>(
      builder: (context, state) {
        return AnimatedPositioned(
          duration: const Duration(milliseconds: 500),
          top: 0,
          left: state,
          right: 0,
          child: AnimatedPadding(
            duration: const Duration(milliseconds: 500),
            padding: EdgeInsets.only(
                top: 40, left: state == 0 ? 100 : 40, right: 40, bottom: 40),
            child: Column(
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Container(
                      decoration: BoxDecoration(
                        color: Colors.grey.withOpacity(0.1),
                        borderRadius: BorderRadius.circular(10),
                      ),
                      width: MediaQuery.of(context).size.width * 0.22,
                      child: ListTile(
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(10),
                        ),
                        leading: Icon(
                          Icons.search_rounded,
                          color: Colors.grey.shade500,
                        ),
                        title: Text(
                          "Search",
                          style: TextStyle(color: Colors.grey.shade500),
                        ),
                        onTap: () {},
                      ),
                    ),
                    const Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        NotificationButton(),
                        SizedBox(width: 20),
                        UserAvatarAndName(),
                      ],
                    )
                  ],
                ),
              ],
            ),
          ),
        );
      },
    );
  }
}
