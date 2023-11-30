import styled from "styled-components";

export const Contain = styled.div`
  border: solid black 1px;
  padding: 10px;
  width: 500px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
`;

export const Description = styled.div``;

export const ButtonWrapper = styled.div`
  display: flex;
  justify-content: space-around;
  width: 60px;
  div {
    &:hover {
      cursor: pointer;
    }
  }
`;

export const DraggableWrapper = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
`;
