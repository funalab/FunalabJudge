import React from "react";
import { UserType } from "../types/UserTypes"; 

export type AuthUserContextType = {
  user: UserType | null;
  signin: (user:UserType, callback:() => void) => void;
  signout: (callback:() => void) => void;
}
const AuthUserContext = React.createContext<AuthUserContextType>({} as AuthUserContextType);

export const useAuthUserContext = ():AuthUserContextType => {
  return React.useContext<AuthUserContextType>(AuthUserContext);
}

type Props = {
  children: React.ReactNode
}

export const AuthUserProvider = (props:Props) => {
  const [user, setUser] = React.useState<UserType | null>(null);

  const signin = (newUser: UserType, callback: () => void) => {
    setUser(newUser);
    callback();
  }

  const signout = (callback: () => void) => {
    setUser(null);
    callback();
  }

  const value:AuthUserContextType = { user, signin, signout };
  return (
    <AuthUserContext.Provider value={value}>
      {props.children}
    </AuthUserContext.Provider>
  );
}