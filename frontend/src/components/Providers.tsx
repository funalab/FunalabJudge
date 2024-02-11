import React from "react";
import { AuthUserProvider } from "../providers/AuthUser";

type Props = {
  children: React.ReactNode
}

export const Providers:React.FC<Props> = (props) => {
  return (
    <>
        {props.children}
    </>
  );
}