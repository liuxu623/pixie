#include <gtest/gtest.h>

#include "src/carnot/funcs/builtins/json_ops.h"
#include "src/carnot/udf/test_utils.h"

namespace pl {
namespace carnot {
namespace builtins {

using types::StringValue;

constexpr char kTestJSONStr[] = R"(
{
  "str_key": {"abc": "def"},
  "int64_key": 34243242341,
  "float64_key": 123423.5234,
  "str_plain": "abc"
})";

TEST(JSONOps, PluckUDF) {
  auto udf_tester = udf::UDFTester<PluckUDF>();
  udf_tester.ForInput(kTestJSONStr, "str_key").Expect(R"({"abc":"def"})");
}

TEST(JSONOps, PluckUDF_plain_string) {
  auto udf_tester = udf::UDFTester<PluckUDF>();
  udf_tester.ForInput(kTestJSONStr, "str_plain").Expect("abc");
}

TEST(JSONOps, PluckUDF_non_existent_key_return_empty) {
  auto udf_tester = udf::UDFTester<PluckUDF>();
  udf_tester.ForInput(kTestJSONStr, "blah").Expect("");
}

TEST(JSONOps, PluckUDF_bad_input_return_empty) {
  auto udf_tester = udf::UDFTester<PluckUDF>();
  udf_tester.ForInput("asdad", "str_key").Expect("");
}

TEST(JSONOps, PluckUDF_non_object_input_return_empty) {
  auto udf_tester = udf::UDFTester<PluckUDF>();
  udf_tester.ForInput("[\"asdad\"]", "str_key").Expect("");
}

TEST(JSONOps, PluckAsInt64UDF) {
  auto udf_tester = udf::UDFTester<PluckAsInt64UDF>();
  udf_tester.ForInput(kTestJSONStr, "int64_key").Expect(34243242341);
}

TEST(JSONOps, PluckAsInt64UDF_bad_input_return_empty) {
  auto udf_tester = udf::UDFTester<PluckAsInt64UDF>();
  udf_tester.ForInput("sdasdsa", "int64_key").Expect(0);
}

TEST(JSONOps, PluckAsIntUDF_non_object_input_return_empty) {
  auto udf_tester = udf::UDFTester<PluckAsInt64UDF>();
  udf_tester.ForInput("[\"asdad\"]", "int64_key").Expect(0);
}

TEST(JSONOps, PluckAsFloat64UDF) {
  auto udf_tester = udf::UDFTester<PluckAsFloat64UDF>();
  udf_tester.ForInput(kTestJSONStr, "float64_key").Expect(123423.5234);
}

TEST(JSONOps, PluckAsFloat64UDF_bad_input_return_empty) {
  auto udf_tester = udf::UDFTester<PluckAsFloat64UDF>();
  udf_tester.ForInput("sdadasd", "float64_key").Expect(0.0);
}

TEST(JSONOps, PluckAsFloatUDF_non_object_input_return_empty) {
  auto udf_tester = udf::UDFTester<PluckAsFloat64UDF>();
  udf_tester.ForInput("[\"asdad\"]", "float64_key").Expect(0.0);
}

TEST(JSONOps, ScriptReferenceUDF_no_args) {
  auto udf_tester = udf::UDFTester<ScriptReferenceUDF<>>();
  auto res = udf_tester.ForInput("text", "px/script").Result();

  rapidjson::Document d;
  d.Parse(res.data());

  EXPECT_TRUE(d.IsObject());
  auto elements = std::distance(d.MemberBegin(), d.MemberEnd());
  EXPECT_EQ(elements, 3);

  EXPECT_TRUE(d["label"].IsString());
  EXPECT_STREQ(d["label"].GetString(), "text");

  EXPECT_TRUE(d["script"].IsString());
  EXPECT_STREQ(d["script"].GetString(), "px/script");

  EXPECT_TRUE(d["args"].IsObject());
  auto args = std::distance(d["args"].MemberBegin(), d["args"].MemberEnd());
  EXPECT_EQ(args, 0);
}

TEST(JSONOps, ScriptReferenceUDF_with_args) {
  auto udf_tester =
      udf::UDFTester<ScriptReferenceUDF<StringValue, StringValue, StringValue, StringValue>>();
  auto res = udf_tester
                 .ForInput("text", "px/script", "my_arg_name", "my_arg_value", "my_arg_name2",
                           "my_arg_value2")
                 .Result();

  rapidjson::Document d;
  d.Parse(res.data());

  EXPECT_TRUE(d.IsObject());
  auto elements = std::distance(d.MemberBegin(), d.MemberEnd());
  EXPECT_EQ(elements, 3);

  EXPECT_TRUE(d["label"].IsString());
  EXPECT_STREQ(d["label"].GetString(), "text");

  EXPECT_TRUE(d["script"].IsString());
  EXPECT_STREQ(d["script"].GetString(), "px/script");

  EXPECT_TRUE(d["args"].IsObject());
  auto args = std::distance(d["args"].MemberBegin(), d["args"].MemberEnd());
  EXPECT_EQ(args, 2);

  EXPECT_TRUE(d["args"]["my_arg_name"].IsString());
  EXPECT_STREQ(d["args"]["my_arg_name"].GetString(), "my_arg_value");

  EXPECT_TRUE(d["args"]["my_arg_name2"].IsString());
  EXPECT_STREQ(d["args"]["my_arg_name2"].GetString(), "my_arg_value2");
}

}  // namespace builtins
}  // namespace carnot
}  // namespace pl
