import Cookies from "./Cookie";
let c = new Cookies();

export const loadUser = () => {
  let jwt = "";
  if ((jwt = c.getCookie("jwt"))) {
    jwt = JSON.parse(jwt);
    let user = {
      id: jwt.id,
      name: jwt.name,
      token: jwt.token
    };
    user.loggedIn = true;
    return user;
  } else {
    return false;
  }
};
