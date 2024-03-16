import React from "react";
import { PageType } from "../types/PageTypes";
import { Navigate, useLocation, useParams } from "react-router-dom";

type Props = {
  component: React.ReactNode;
  pageType: PageType,
}

export const RouteAuthGuard: React.FC<Props> = (props) => {
  let allowRoute = false;
  let message = "";
  const location = useLocation();

  const authUserName = localStorage.getItem("authUserName");
  const authJoinedDate = localStorage.getItem("authJoinedDate");
  const authUserExp = localStorage.getItem("authUserExp");

  if ( authUserName && authJoinedDate && authUserExp ) {
    if (Number(authUserExp) < Date.now() / 1000) {
      allowRoute = false;
      message = "ログイン保持期限が切れました。再度ログインしてください。";
    } else if (props.pageType === PageType.Public) {
      allowRoute = true;
    } else if (props.pageType === PageType.Private) {
      [allowRoute, message] = CheckAccessPermission({
        authUserName: authUserName,
        authJoinedDate: new Date(authJoinedDate)
      });
    } else {
      // unknown page type
    }
  } else {
    allowRoute = false;
    message = "コンテンツの閲覧にはログインが必要です。";
  }

  if (!allowRoute) {
    alert(message);
    return <Navigate to="/login" state={{from: location}} replace={false} />
  }

  return <>{props.component}</>;

}

type AuthUserProps = {
  authUserName: string,
  authJoinedDate: Date
}

export const CheckAccessPermission = (props:AuthUserProps): [boolean, string] => {
    const { userName } = useParams()
    if (props.authJoinedDate.getFullYear() < new Date().getFullYear()) {
        return [true, ""];  // senior student
    } else {
        if (props.authUserName == userName) {
            return [true, ""];  // matched userName
        } else {
            return [false, "ページへのアクセス権がありません。アクセス権のあるアカウントでログインしてください。"];  // unmatched userName
        }
    }
}