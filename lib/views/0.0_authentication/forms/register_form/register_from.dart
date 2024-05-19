import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/utils/custom_text_field.dart';
import 'package:flutter/material.dart';

class RegisterForm extends StatelessWidget {
  const RegisterForm({super.key});

  @override
  Widget build(BuildContext context) {
    return Builder(builder: (context) {
      final size = MediaQuery.of(context).size;
      final registerBloc = context.read<RegisterBloc>();
      return Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 100),
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                "Create new Account.",
                style: ThemeX.titleText.copyWith(
                  fontSize: size.width * 0.039,
                ),
              ),
              const SizedBox(height: 10),
              Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text("Already A Member?",
                      style: TextStyle(
                        color: Colors.grey.shade600,
                      )),
                  TextButton(
                      onPressed: () =>
                          context.read<AuthSwapperBloc>().swap(context),
                      child: const Text("Log In")),
                ],
              ),
            ],
          ),
          const SizedBox(height: 20),
          Row(
            children: List.generate(
              2,
              (index) => Expanded(
                child: CustomTextField(
                  keyboardType: TextInputType.text,
                  controller: index == 0
                      ? registerBloc.firstNameController
                      : registerBloc.lastNameController,
                  formKey: index == 0 ? registerBloc.userNameFromKey : null,
                  validator: index == 0 ? registerBloc.userNameValidator : null,
                  label: Text(
                    index == 0 ? "First Name" : "Last Name",
                    style: const TextStyle(color: Colors.white),
                  ),
                  // prefixIcon: const Icon(Icons.person, color: Colors.white),
                  suffixIcon: const Icon(Icons.assignment_ind_rounded,
                      color: Colors.white),
                ),
              ),
            ),
          ),
          const SizedBox(height: 10),
          CustomTextField(
            keyboardType: TextInputType.emailAddress,
            formKey: registerBloc.emailFromKey,
            validator: registerBloc.emailValidator,
            controller: registerBloc.emailController,
            label: const Text(
              "Email",
              style: TextStyle(color: Colors.white),
            ),
            suffixIcon: const Icon(Icons.email, color: Colors.white),
          ),
          const SizedBox(height: 10),
          CustomTextField(
            keyboardType: TextInputType.text,
            formKey: registerBloc.passwordFromKey,
            validator: registerBloc.passwordValidator,
            controller: registerBloc.passwordController,
            label: const Text(
              "Password",
              style: TextStyle(color: Colors.white),
            ),
            prefixIcon: const Icon(Icons.lock, color: Colors.white),
          ),
          const SizedBox(height: 40),
          Center(
            child: ValueListenableBuilder(
                valueListenable: registerBloc.isLoading,
                builder: (_, isLoading, __) {
                  return ElevatedButton(
                    onPressed: !isLoading
                        ? () => registerBloc.register(context)
                        : null,
                    child: Padding(
                      padding: const EdgeInsets.symmetric(
                          vertical: 20, horizontal: 40),
                      child: AnimatedSwitcher(
                        duration: const Duration(milliseconds: 300),
                        child: isLoading
                            ? const SizedBox(
                                width: 20,
                                height: 20,
                                child: CircularProgressIndicator(),
                              )
                            : const Text(
                                "Create Account",
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                      ),
                    ),
                  );
                }),
            // Row(
            //   children: [
            //     Expanded(
            //       child: Container(
            //         height: 1,
            //         color: Colors.grey.shade600,
            //       ),
            //     ),
            //     Padding(
            //       padding:
            //           const EdgeInsets.symmetric(horizontal: 8, vertical: 20),
            //       child: Text(
            //         "Or",
            //         style: TextStyle(
            //           color: Colors.grey.shade600,
            //         ),
            //       ),
            //     ),
            //     Expanded(
            //       child: Container(
            //         height: 1,
            //         color: Colors.grey.shade600,
            //       ),
            //     ),
            //   ],
            // ),
            // const SizedBox(height: 20),
            // Center(
            //   child: IconButton(
            //     onPressed: () {},
            //     icon: const Icon(Icons.g_mobiledata, color: Colors.red),
            //   ),
          ),
        ],
      );
    });
  }
}
