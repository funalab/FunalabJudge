import React from "react";
import { PageType } from "../types/PageTypes";
import { useAuthUserContext } from "../providers/AuthUser"
import { AuthPageUser } from "./AuthPageUser";
import { Navigate, useLocation } from "react-router-dom";

type Props = {
  component: React.ReactNode;
  redirect: string,
  pageType: PageType,
}

export const RouteAuthGuard: React.FC<Props> = (props) => {
  const authUser = useAuthUserContext().user;

  let allowRoute = false;
  if ( authUser ) {
    allowRoute = AuthPageUser({pageType: props.pageType, authUser: authUser});
  }

  if (!allowRoute) {
    return <Navigate to={props.redirect} state={{from:useLocation()}} replace={false} />
  }

  return <>{props.component}</>;

}