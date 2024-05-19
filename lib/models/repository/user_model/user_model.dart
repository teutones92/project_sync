// ignore_for_file: public_member_api_docs, sort_constructors_first
class UserModel {
  final int id;
  String username;
  String email;
  String dob;
  String phoneNumber;
  String countryCode;
  String countryPhoneCode;
  String langCode;
  String passwordHash;
  String password;
  String userAvatarPath;
  bool darkMode;
  UserModel({
    required this.id,
    required this.username,
    required this.email,
    required this.dob,
    required this.phoneNumber,
    required this.countryCode,
    required this.countryPhoneCode,
    required this.langCode,
    required this.passwordHash,
    required this.password,
    required this.userAvatarPath,
    required this.darkMode,
  });

  factory UserModel.fromJson(Map<String, Object?> json) {
    return UserModel(
        id: json['id'] as int,
        username: json['username'] as String,
        email: json['email'] as String,
        dob: json['DoB'] as String,
        phoneNumber: json['phone_number'] as String,
        countryCode: json['country_code'] as String,
        countryPhoneCode: json['country_phone_code'] as String,
        langCode: json['lang_code'] as String,
        passwordHash: json['password_hash'] as String,
        password: json['password'] as String,
        userAvatarPath: json['user_avatar_path'] as String,
        darkMode: json['dark_mode'] as bool);
  }

  Map<String, Object?> toJson() {
    return {
      'id': id,
      'username': username,
      'email': email,
      'DoB': dob,
      'phone_number': phoneNumber,
      'country_code': countryCode,
      'country_phone_code': countryPhoneCode,
      'lang_code': langCode,
      'password_hash': passwordHash,
      'password': password,
      'user_avatar_path': userAvatarPath,
      'dark_mode': darkMode
    };
  }
}
