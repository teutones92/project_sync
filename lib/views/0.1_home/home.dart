import 'package:project_sync/core/config/bloc_config.dart';
import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/views/0.1_home/widgets/landing/landing.dart';
import 'package:project_sync/views/0.8_create_project/create_project.dart';
import '../0.8.1_create_contact/create_contact.dart';
import 'widgets/home_header/home_header.dart';
import 'widgets/side_menu/side_menu.dart';
import 'widgets/toggle_menu_button/toggle_menu_button.dart';

class Home extends StatelessWidget {
  const Home({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: BlocBuilder<SideMenuBloc, int>(
        builder: (context, state) {
          return FloatingActionButton(
            onPressed: () {
              switch (state) {
                case 0:
                  NavigatorX.push(
                      context: context, page: const CreateProject());
                  break;
                case 5:
                  NavigatorX.push(
                      page: const CreateContact(), context: context);
                  break;
                default:
              }
            },
            child: AnimatedSwitcher(
              duration: const Duration(milliseconds: 500),
              child: showIcon(state),
            ),
          );
        },
      ),
      body: Stack(
        children: [
          BlocBuilder<AnimateSideMenuBloc, double>(
            builder: (context, state) {
              return AnimatedPositioned(
                duration: const Duration(milliseconds: 500),
                top: 120,
                left: state,
                child: Container(
                  width: MediaQuery.of(context).size.width,
                  height: 1,
                  decoration: BoxDecoration(
                    color: Theme.of(context).brightness == Brightness.light
                        ? Colors.grey.shade300
                        : Colors.grey.shade700,
                  ),
                ),
              );
            },
          ),
          const SideMenu(),
          const HomeHeader(),
          const ToggleMenuButton(),
          const LandingWidget(),
        ],
      ),
    );
  }

  Icon showIcon(int state) {
    switch (state) {
      case 0:
        return Icon(key: ValueKey<int>(state), Icons.add_card_rounded);
      case 5:
        return Icon(key: ValueKey<int>(state), Icons.person_add);
      default:
        return Icon(key: ValueKey<int>(state), Icons.add);
    }
  }
}
