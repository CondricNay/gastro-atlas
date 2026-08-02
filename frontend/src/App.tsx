import { BrowserRouter, Routes, Route } from "react-router-dom";
import IngredientPage from "./pages/IngredientPage";


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route 
          path="/ingredients/:slug" element={<IngredientPage />}
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;