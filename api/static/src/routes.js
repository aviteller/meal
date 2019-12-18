import Cookies from "./helpers/Cookie";
import Home from "./User/components/Home.svelte";
import Login from "./Auth/components/Login.svelte";
import Register from "./Auth/components/Register.svelte";
import Meal from "./Meal/Meal.svelte";

let c = new Cookies();

const userLoggedIn = () => {
  let jwt;
  if ((jwt = c.getCookie("jwt"))) {
    return true;
  } else return false;
};

let routes;
const urlParams = new URLSearchParams(window.location.search);
if (!urlParams.has("routemap")) {
  routes = {
    "/": userLoggedIn() ? Home : Login,
    "/Login": Login,
    "/Register": Register,
    "/meal/:id": userLoggedIn() ? Meal : Login
  };
}

export default routes;
