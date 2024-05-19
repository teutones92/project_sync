import 'package:project_sync/core/extensions.dart';
import 'package:flutter/material.dart';

class CustomTextField extends StatelessWidget {
  const CustomTextField(
      {super.key,
      this.hintText,
      this.controller,
      this.obscureText,
      this.label,
      this.keyboardType,
      this.suffixIcon,
      this.prefixIcon,
      this.validator,
      this.formKey,
      this.onChanged});
  final String? hintText;
  final Widget? label;
  final TextEditingController? controller;
  final bool? obscureText;
  final TextInputType? keyboardType;
  final Widget? suffixIcon;
  final Widget? prefixIcon;
  final String? Function(String?)? validator;
  final GlobalKey<FormState>? formKey;
  final void Function(String)? onChanged;

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 5,
      color: ThemeX.darkBlueColor,
      child: Padding(
        padding: const EdgeInsets.all(4),
        child: Form(
          key: formKey,
          child: TextFormField(
            controller: controller,
            validator: validator,
            style: const TextStyle(color: Colors.white),
            decoration: InputDecoration(
              label: label,
              labelStyle: const TextStyle(color: Colors.white),
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(10),
              ),
              prefixIcon: prefixIcon,
              suffixIcon: suffixIcon,
            ),
          ),
        ),
      ),
    );
  }
}
