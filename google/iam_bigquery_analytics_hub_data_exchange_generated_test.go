// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccBigqueryAnalyticsHubDataExchangeIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryAnalyticsHubDataExchangeIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_analytics_hub_data_exchange_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s roles/viewer", acctest.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf_test_my_data_exchange%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccBigqueryAnalyticsHubDataExchangeIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_analytics_hub_data_exchange_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s roles/viewer", acctest.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf_test_my_data_exchange%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigqueryAnalyticsHubDataExchangeIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccBigqueryAnalyticsHubDataExchangeIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_analytics_hub_data_exchange_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s roles/viewer user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf_test_my_data_exchange%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigqueryAnalyticsHubDataExchangeIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryAnalyticsHubDataExchangeIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_analytics_hub_data_exchange_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", acctest.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf_test_my_data_exchange%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigqueryAnalyticsHubDataExchangeIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_bigquery_analytics_hub_data_exchange_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", acctest.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf_test_my_data_exchange%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigqueryAnalyticsHubDataExchangeIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_analytics_hub_data_exchange" "data_exchange" {
  location         = "US"
  data_exchange_id = "tf_test_my_data_exchange%{random_suffix}"
  display_name     = "tf_test_my_data_exchange%{random_suffix}"
  description      = "example data exchange%{random_suffix}"
}

resource "google_bigquery_analytics_hub_data_exchange_iam_member" "foo" {
  project = google_bigquery_analytics_hub_data_exchange.data_exchange.project
  location = google_bigquery_analytics_hub_data_exchange.data_exchange.location
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.data_exchange.data_exchange_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccBigqueryAnalyticsHubDataExchangeIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_analytics_hub_data_exchange" "data_exchange" {
  location         = "US"
  data_exchange_id = "tf_test_my_data_exchange%{random_suffix}"
  display_name     = "tf_test_my_data_exchange%{random_suffix}"
  description      = "example data exchange%{random_suffix}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_bigquery_analytics_hub_data_exchange_iam_policy" "foo" {
  project = google_bigquery_analytics_hub_data_exchange.data_exchange.project
  location = google_bigquery_analytics_hub_data_exchange.data_exchange.location
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.data_exchange.data_exchange_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigqueryAnalyticsHubDataExchangeIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_analytics_hub_data_exchange" "data_exchange" {
  location         = "US"
  data_exchange_id = "tf_test_my_data_exchange%{random_suffix}"
  display_name     = "tf_test_my_data_exchange%{random_suffix}"
  description      = "example data exchange%{random_suffix}"
}

data "google_iam_policy" "foo" {
}

resource "google_bigquery_analytics_hub_data_exchange_iam_policy" "foo" {
  project = google_bigquery_analytics_hub_data_exchange.data_exchange.project
  location = google_bigquery_analytics_hub_data_exchange.data_exchange.location
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.data_exchange.data_exchange_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigqueryAnalyticsHubDataExchangeIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_analytics_hub_data_exchange" "data_exchange" {
  location         = "US"
  data_exchange_id = "tf_test_my_data_exchange%{random_suffix}"
  display_name     = "tf_test_my_data_exchange%{random_suffix}"
  description      = "example data exchange%{random_suffix}"
}

resource "google_bigquery_analytics_hub_data_exchange_iam_binding" "foo" {
  project = google_bigquery_analytics_hub_data_exchange.data_exchange.project
  location = google_bigquery_analytics_hub_data_exchange.data_exchange.location
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.data_exchange.data_exchange_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccBigqueryAnalyticsHubDataExchangeIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_analytics_hub_data_exchange" "data_exchange" {
  location         = "US"
  data_exchange_id = "tf_test_my_data_exchange%{random_suffix}"
  display_name     = "tf_test_my_data_exchange%{random_suffix}"
  description      = "example data exchange%{random_suffix}"
}

resource "google_bigquery_analytics_hub_data_exchange_iam_binding" "foo" {
  project = google_bigquery_analytics_hub_data_exchange.data_exchange.project
  location = google_bigquery_analytics_hub_data_exchange.data_exchange.location
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.data_exchange.data_exchange_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
