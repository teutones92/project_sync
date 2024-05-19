import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/views/0.0_authentication/forms/login_form/login_form.dart';
import 'package:flutter/material.dart';

import 'widgets/bg_widget/bg_widget.dart';
import 'widgets/header_widget/header_widget.dart';
import 'forms/register_form/register_from.dart';

class Authentication extends StatelessWidget {
  const Authentication({super.key});

  @override
  Widget build(BuildContext context) {
    final size = MediaQuery.of(context).size;
    double widthMin = 700;
    return Scaffold(
      body: Stack(
        fit: StackFit.expand,
        children: [
          const BackGroundWidget(),
          const HeaderWidget(),
          Positioned(
            left: size.width <= widthMin ? 0 : 120,
            height: size.height,
            width: size.width <= widthMin ? size.width : size.width * 0.40,
            child: Padding(
              padding: EdgeInsets.symmetric(
                  horizontal: size.width <= widthMin ? 8 : 0),
              child: BlocBuilder<AuthSwapperBloc, bool>(
                builder: (context, state) {
                  return AnimatedSwitcher(
                    duration: const Duration(milliseconds: 800),
                    switchInCurve: Curves.easeIn,
                    switchOutCurve: Curves.easeOut,
                    child: !state ? const RegisterForm() : const LoginForm(),
                  );
                },
              ),
            ),
          ),
        ],
      ),
    );
  }
}
