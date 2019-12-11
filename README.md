# goleekwars

goleekwars is a Go client library for accessing the [leekwars API](https://leekwars.com)

## Installation
```shell
go get github.com/punkymaniac/goleekwars
```

## Quick start
```go
import "github.com/punkymaniac/goleekwars"

apileek := apileekwars.NewApi()
err := api.Auth("username", "password")
if err != nil {
    panic(err)
}

// Get all ais's farmer
ais, err := api.Ai.GetFarmerAis()
if err != nil {
    panic(err)
}

aiFiles := ais.Ais
aiFolders := ais.Folders

// ...

```

## Notes
Current Api call implemented:
* [ ] /ai/change-folder/
* [ ] /ai/delete/
* [ ] /ai/get/
* [ ] /ai/get-farmer-ais/
* [ ] /ai/new/
* [ ] /ai/rename/
* [ ] /ai/save/
* [ ] /ai/test/
* [ ] /ai/test-new/
* [ ] /ai/test-v2/
* [ ] /ai-folder/change-folder/
* [ ] /ai-folder/delete/
* [ ] /ai-folder/new/
* [ ] /ai-folder/rename
* [ ] /changelog/get/
* [ ] /changelog/get-last/
* [ ] /chip/get-all/
* [ ] /chip/get-templates/
* [ ] /constant/get-all/
* [ ] /country/get-all/
* [ ] /farmer/accept-terms/
* [ ] /farmer/activate/
* [ ] /farmer/change-country/
* [ ] /farmer/change-email1/
* [ ] /farmer/change-email2/
* [ ] /farmer/change-email3/
* [ ] /farmer/chnage-password/
* [ ] /farmer/disconnect/
* [ ] /farmer/get/
* [ ] /farmer/get-connected/
* [ ] /farmer/get-from-token/
* [ ] /farmer/login/
* [ ] /farmer/login-token/
* [ ] /farmer/register/
* [ ] /farmer/register-tournament/
* [ ] /farmer/rich-tooltip/
* [ ] /farmer/set-avatar/
* [ ] /farmer/set-github/
* [ ] /farmer/set-in-garden/
* [ ] /farmer/set-website/
* [ ] /farmer/unregister/
* [ ] /farmer/unregister-tournament/
* [ ] /farmer/update/
* [ ] /fight/comment/
* [ ] /fight/get/
* [ ] /fight/get-logs/
* [ ] /forum/search2/
* [ ] /function/get-all/
* [ ] /function/get-categories/
* [ ] /function/operations/
* [ ] /garden/get/
* [ ] /garden/get-composition-opponents/
* [ ] /garden/get-farmer-challenge/
* [ ] /garden/get-farmer-opponents/
* [ ] /garden/get-leek-opponents/
* [ ] /garden/get-solo-challenge/
* [ ] /garden/start-farmer-challenge/
* [ ] /garden/start-farmer-fight/
* [ ] /garden/start-solo-challenge/
* [ ] /garden/start-solo-fight/
* [ ] /garden/start-team-fight/
* [ ] /hat/get-all/
* [ ] /hat/get-templates/
* [ ] /history/get-farmer-history/
* [ ] /history/get-leek-history/
* [ ] /lang/get/
* [ ] /leek/add-chip/
* [ ] /leek/add-weapon/
* [ ] /leek/create/
* [ ] /leek/delete-register/
* [ ] /leek/get/
* [ ] /leek/get-count/
* [ ] /leek/get-image/
* [ ] /leek/get-level-popup/
* [ ] /leek/get-next-price/
* [ ] /leek/get-private/
* [ ] /leek/get-registers/
* [ ] /leek/register-tournament/
* [ ] /leek/remove-ai/
* [ ] /leek/remove-chip/
* [ ] /leek/remove-hat/
* [ ] /leek/remove-weapon/
* [ ] /leek/rename-crystals/
* [ ] /leek/rename-habs/
* [ ] /leek/rich-tooltip/
* [ ] /leek/set-ai/
* [ ] /leek/set-hat/
* [ ] /leek/set-in-garden/
* [ ] /leek/set-popup-level-seen/
* [ ] /leek/set-register/
* [ ] /leek/spend-capital/
* [ ] /leek/unregister-tournament/
* [ ] /leek/use-potions/
* [ ] /leek-wars/version/
* [ ] /market/buy-crystals/
* [ ] /market/buy-habs/
* [ ] /market/get-item-templates/
* [ ] /market/sell-habs/
* [ ] /message/create-conversation/
* [ ] /message/find-conversation/
* [ ] /message/get-latest-conversations/
* [ ] /message/get-messages/
* [ ] /message/quit-conversation/
* [ ] /message/send-message/
* [ ] /notification/get-latest/
* [ ] /notification/read/
* [ ] /notification/read-all/
* [ ] /potion/get-all/
* [ ] /ranking/fun/
* [ ] /ranking/get/
* [ ] /ranking/get-active/
* [ ] /ranking/get-farmer-rank/
* [ ] /ranking/get-farmer-rank-active/
* [ ] /ranking/get-home-ranking/
* [ ] /ranking/get-leek-rank/
* [ ] /ranking/get-leek-rank-active/
* [ ] /ranking/get-team-rank/
* [ ] /ranking/get-team-rank-active/
* [ ] /ranking/search/
* [ ] /service/get-all/
* [ ] /summon/get-templates/
* [ ] /team/accept-candidacy/
* [ ] /team/ban/
* [ ] /team/cancel-candidacy/
* [ ] /team/cancel-candidacy-for-team/
* [ ] /team/change-description/
* [ ] /team/change-member-grade/
* [ ] /team/change-owner/
* [ ] /team/create/
* [ ] /team/create-composition/
* [ ] /team/delete-composition/
* [ ] /team/dissolve/
* [ ] /team/get/
* [ ] /team/get-connected/
* [ ] /team/get-private/
* [ ] /team/move-leek/
* [ ] /team/quit/
* [ ] /team/register-tournament/
* [ ] /team/reject-condidacy/
* [ ] /team/send-candidacy/
* [ ] /team/set-emblem/
* [ ] /team/set-opened/
* [ ] /team/unregister-tournament/
* [ ] /test-leek/delete/
* [ ] /test-leek/get-all/
* [ ] /test-leek/new/
* [ ] /test-leek/update/
* [ ] /test-map/delete/
* [ ] /test-map/get-all/
* [ ] /test-map/new/
* [ ] /test-map/update/
* [ ] /test-scenario/delete/
* [ ] /test-scenario/get-all/
* [ ] /test-scenario/new/
* [ ] /test-scenario/update/
* [ ] /tournament/comment/
* [ ] /tournament/get/
* [ ] /trophy/get-admin/
* [ ] /trophy/get-all/
* [ ] /trophy/get-categories/
* [ ] /trophy/get-farmer-trophies/
* [ ] /trophy/unlock/
* [ ] /weapon/get-all/
* [ ] /weapon/get-templates/

