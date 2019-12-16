<script>
  import Header from "./UI/Header.svelte";
  import routes from "./routes";
  import Router, { location, querystring } from "svelte-spa-router";

	let title = "App";

</script>

<style>
  main {
    margin-top: 5rem;
  }
</style>

<Header>{title}</Header>
<main>
  <Router {routes} />
</main>
