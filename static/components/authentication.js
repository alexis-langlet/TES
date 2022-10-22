app.component("Authentication", {
  data() {
    return {
      authentication: "credentials",
    };
  },
  updated() {
    window.componentHandler.upgradeElements(this.$el)
  },
  template: `<div>
  <div>
    <input type="radio" id="token_auth" value="token_auth" v-model="authentication" />
    <label for="token_auth">api token</label>

    <input
      type="radio"
      id="credentials"
      value="credentials"
      v-model="authentication"
      checked
    />
    <label for="credentials">username/password</label>
  </div>
  <div v-if="authentication === 'credentials'">
    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
      <input class="mdl-textfield__input" type="text" id="username" required />
      <label class="mdl-textfield__label" for="username">username</label>
    </div>

    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
      <input
        class="mdl-textfield__input"
        type="password"
        id="password"
        required
      />
      <label class="mdl-textfield__label" for="password">password</label>
    </div>
  </div>
  <div v-else>
    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
      <input class="mdl-textfield__input" type="text" id="token" required />
      <label class="mdl-textfield__label" for="token">Api token</label>
    </div>
  </div>
</div>
`,
});
