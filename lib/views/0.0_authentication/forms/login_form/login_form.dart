import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/utils/custom_text_field.dart';
import 'package:flutter/material.dart';

import '../../../../core/config/bloc_config.dart';

class LoginForm extends StatelessWidget {
  const LoginForm({super.key});

  @override
  Widget build(BuildContext context) {
    return Builder(builder: (context) {
      final size = MediaQuery.of(context).size;
      final loginBloc = context.read<LoginBloc>();
      return Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 100),
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                "Welcome Back.",
                style: ThemeX.titleText.copyWith(
                  fontSize: size.width * 0.039,
                ),
              ),
              Wrap(
                // mainAxisSize: MainAxisSize.min,
                crossAxisAlignment: WrapCrossAlignment.center,
                children: [
                  Padding(
                    padding: const EdgeInsets.only(right: 4),
                    child: Text("Don't Have An Account?",
                        style: TextStyle(
                          color: Colors.grey.shade600,
                        )),
                  ),
                  TextButton(
                      style: ButtonStyle(
                        padding: WidgetStateProperty.all(
                          EdgeInsets.zero,
                        ),
                      ),
                      onPressed: () =>
                          context.read<AuthSwapperBloc>().swap(context),
                      child: const Text("Sign Up")),
                ],
              ),
            ],
          ),
          const SizedBox(height: 20),
          CustomTextField(
            formKey: loginBloc.emailFormKey,
            keyboardType: TextInputType.emailAddress,
            controller: loginBloc.emailController,
            validator: loginBloc.emailValidator,
            label: const Text(
              "Email",
              style: TextStyle(color: Colors.white),
            ),
            prefixIcon: const Icon(Icons.email, color: Colors.white),
          ),
          const SizedBox(height: 10),
          CustomTextField(
            formKey: loginBloc.passwordFormKey,
            keyboardType: TextInputType.visiblePassword,
            controller: loginBloc.passwordController,
            validator: loginBloc.passwordValidator,
            label: const Text(
              "Password",
              style: TextStyle(color: Colors.white),
            ),
            prefixIcon: const Icon(Icons.lock, color: Colors.white),
            suffixIcon: const Icon(Icons.remove_red_eye, color: Colors.white),
          ),
          const SizedBox(height: 10),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              TextButton(
                onPressed: () {
                  showDialog(
                    context: context,
                    builder: (context) => AlertDialog(
                      title: const Text("Forgot Password"),
                      content: const Text(
                        "We sent an email with a link to reset your password. Please check your email.",
                      ),
                      actions: [
                        TextButton(
                          onPressed: () => Navigator.pop(context),
                          child: const Text("Close"),
                        ),
                      ],
                    ),
                  );
                },
                child: Text(
                  "Forgot Password?",
                  style: TextStyle(color: Colors.grey.shade600),
                ),
              ),
            ],
          ),
          const SizedBox(height: 40),
          Center(
            child: ValueListenableBuilder(
                valueListenable: loginBloc.isLoading,
                builder: (_, isLoading, __) {
                  return ElevatedButton(
                    onPressed: isLoading
                        ? null
                        : () => loginBloc.login(context, fromLogin: true),
                    child: Padding(
                      padding: const EdgeInsets.symmetric(
                          vertical: 20, horizontal: 40),
                      child: AnimatedSwitcher(
                        duration: const Duration(milliseconds: 800),
                        child: isLoading
                            ? const CircularProgressIndicator()
                            : const Text("Log In"),
                      ),
                    ),
                  );
                }),
          ),
        ],
      );
    });
  }
}
