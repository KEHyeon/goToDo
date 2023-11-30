import styled from "styled-components";

export const Contain = styled.div`
  border: solid black 1px;
  padding: 10px;
  width: 500px;
  min-height: 700px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  position: relative;
  padding-top: 70px;
`;

export const ButtonWrapper = styled.div`
  display: flex;
  justify-content: space-around;
  width: 60px;
  div {
    &:hover {
      cursor: pointer;
    }
  }
  position: relative;
`;
export const DateButtonWrapper = styled.div`
  display: flex;
  justify-content: space-around;
  width: 100px;
  div {
    &:hover {
      cursor: pointer;
    }
  }
`;

export const DateWrapper = styled.div`
  display: flex;
  flex-direction: column;
  border: solid 1px;
  align-items: center;
  position: absolute;
  top: 20px;
`;

export const DraggableWrapper = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
`;
