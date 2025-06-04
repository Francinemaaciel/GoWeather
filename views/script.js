const form = document.getElementById('formclima');
const resultado = document.getElementById('resultado');

form.addEventListener('submit', async (e) => {
    e.preventDefault();
    resultado.innerHTML = 'Carregando...';

    const cidade = document.getElementById('cidade').value;

    try {
        const res = await fetch(`/clima/${cidade}`);
        const data = await res.json();

        if (data.erro) {
            resultado.innerHTML = `<p class="erro">Erro: ${data.erro}</p>`;
        } else {
            const temp = Math.round(data.main.temp);
            const feelsLike = Math.round(data.main.feels_like);
            const condicao = data.weather[0].description;
            const nome = data.name;

            resultado.innerHTML = `
                <h3>Clima em ${nome}</h3>
                <p>Temperatura: ${temp}°C</p>
                <p>Sensação térmica: ${feelsLike}°C</p>
                <p>Condição: ${condicao}</p>
            `;
        }
    } catch (erro) {
        resultado.innerHTML = `<p class="erro">Erro ao conectar com o servidor</p>`;
    }
});
