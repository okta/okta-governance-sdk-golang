# Changelog
All notable changes to this project will be documented in this file.

## v2.0.0

Regenerated the SDK from an updated OpenAPI spec. This release includes breaking changes — see below. Thanks [@pranav-okta](https://github.com/pranav-okta)

### Breaking changes

- `ResourceSettingsMutable.Type` changed from `CampaignResourceType` to `*CampaignResourceType`; `NewResourceSettingsMutable` no longer takes a `type` argument.
- `OrgRequestSettingsPatchable.SubprocessorsAcknowledged` changed from `bool` to `*bool`; `NewOrgRequestSettingsPatchable` no longer takes a `subprocessorsAcknowledged` argument.
- `CampaignLinks.EndCampaign` and `CampaignLinks.LaunchCampaign` changed from `Link` to `*Link`; `NewCampaignLinks` now takes 2 arguments instead of 4.
- `ReviewLinks.ReassignReview` changed from `Link` to `*Link`; `NewReviewLinks` now takes 1 argument instead of 2.
- `NewReviewSparse` parameter order changed (the `ReviewLinks` argument moved from 1st to 6th position).
- `NewEntitlementValue2`, `NewEntitlementValueWithParent`, `NewEntitlementsFullWithParent`, and `NewEntitlementsListObject` each gained 4 new required trailing arguments for audit metadata (`createdBy`, `created`, `lastUpdated`, `lastUpdatedBy`).
- Removed `MyAccessCertificationReviewsAPI.ListMyManagedConnections`; replaced by `MyAccessCertificationReviewsAPI.ListMyResourceConnections`.
- Removed `ApiReplaceEntitlementRequest.EntitlementsFullWithParent`; replaced by `ApiReplaceEntitlementRequest.EntitlementUpdatable`.
- Removed `Integrations.Data` field and its accessor methods.
- Removed the following types, superseded by newer resource/connection models: `CriteriaValue`, `CriteriaValueCreatable`, `EntitlementBundlesList`, `GovernanceLabel`, `GroupsArrayCreatableInner`, `Integration`, `LinkNext`, `ManagedConnectionAppInstance`, `ManagedConnectionServiceAccount`, `ManagedConnectionVaultedSecret`, `MyManagedConnection`, `MyManagedConnectionCommon`, `MyManagedConnections`, `MyManagedConnectionsLinks` (and their `Nullable*`/`New*` helpers).

### Additions

- New API groups: `CollectionsV2API`, `MyTasksAPI`, `OperationsAPI`, `TasksAPI`.
- New methods on `MyAccessCertificationReviewsAPI` (bulk campaign decisions, campaign reviews, resource connections) and `OrgGovernanceSettingsAPI` (org integrations, certification settings).
- New model families: entitlement drift tracking, AI-agent-connected resources, resource assets/inventory, bulk review, collections v2, tasks.
- Bumped `github.com/okta/okta-sdk-golang/v6` dependency from v6.0.2 to v6.1.7.

## v1.1.0

- Update the `openapi` spec and regenerate the SDK. Thanks [@pranav-okta](https://github.com/pranav-okta)

## v1.0.1

- Fixes bug in Get and List Entitlement Bundles APIs. Thanks [@aditya-okta](https://github.com/aditya-okta).

## v1.0.0

- Initial release of the SDK for managing the Okta Identity Governance API. Thanks [@aditya-okta](https://github.com/aditya-okta).
