import { styled } from "styled-components";

export const Contain = styled.div`
  width: 80%;
  height: 300px;
  color: black;
  border: solid black 1px;
  margin: 10px;
  padding: 10px;
  justify-content: space-around;
  display: flex;
  flex-direction: column;
  align-items: center;
`;

export const Form = styled.div`
  border: solid black 1px;
  width: 80%;
  height: 80%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  div {
    margin-top: 20px;
    width: 100%;
  }
`;
export const Input = styled.input`
  display: block;
  height: 40px;
  width: 95%;
`;

export const Button = styled.button`
  width: 150px;
  height: 50px;
  margin-top: 50px;
`;
