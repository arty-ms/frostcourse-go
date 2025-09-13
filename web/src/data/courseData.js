// Список всех доступных курсов
export const courses = [
    {
        id: 'js-basics',
        title: 'Основы JavaScript',
        description: 'Изучите основы программирования на JavaScript: переменные, функции, массивы и объекты.',
        difficulty: 'Начинающий',
        duration: '30 мин',
        steps: 10,
        icon: '🟨',
        color: '#f7df1e'
    },
    {
        id: 'react-intro',
        title: 'Введение в React',
        description: 'Познакомьтесь с основами React: компоненты, состояние, события.',
        difficulty: 'Средний',
        duration: '45 мин',
        steps: 8,
        icon: '⚛️',
        color: '#61dafb',
        comingSoon: true
    },
    {
        id: 'node-basics',
        title: 'Node.js для начинающих',
        description: 'Серверное программирование на JavaScript с Node.js.',
        difficulty: 'Средний',
        duration: '40 мин',
        steps: 9,
        icon: '🟢',
        color: '#339933',
        comingSoon: true
    },
    {
        id: 'typescript-intro',
        title: 'Основы TypeScript',
        description: 'Типизированный JavaScript для больших проектов.',
        difficulty: 'Продвинутый',
        duration: '50 мин',
        steps: 12,
        icon: '🔷',
        color: '#3178c6',
        comingSoon: true
    }
];

// Данные курса JavaScript
export const courseSteps = [
    {
        id: 1,
        title: "Переменные",
        theory: "Переменные в JavaScript используются для хранения данных. Есть три способа объявления переменных: var, let и const.",
        code: `let name = 'Артем';
const age = 25;
var city = 'Москва';

console.log(name); // Артем
console.log(age);  // 25`,
        quiz: {
            question: "Какое ключевое слово используется для объявления константы?",
            options: ["var", "let", "const", "function"],
            correct: 2
        }
    },
    {
        id: 2,
        title: "Типы данных",
        theory: "JavaScript имеет несколько примитивных типов данных: string (строка), number (число), boolean (логический), undefined, null.",
        code: `let str = "Привет мир";     // string
let num = 42;             // number
let bool = true;          // boolean
let empty;                // undefined
let nothing = null;       // null

console.log(typeof str);  // "string"`,
        quiz: {
            question: "Какой тип данных у переменной: let x = 3.14;",
            options: ["string", "number", "float", "decimal"],
            correct: 1
        }
    },
    {
        id: 3,
        title: "Функции",
        theory: "Функции - это блоки кода, которые можно многократно использовать. Функции принимают параметры и могут возвращать значения.",
        code: `function greet(name) {
    return "Привет, " + name + "!";
}

// Стрелочная функция
const add = (a, b) => {
    return a + b;
};

console.log(greet("Анна")); // Привет, Анна!
console.log(add(5, 3));     // 8`,
        quiz: {
            question: "Как объявить стрелочную функцию?",
            options: ["function() {}", "() => {}", "def function()", "func() {}"],
            correct: 1
        }
    },
    {
        id: 4,
        title: "Массивы",
        theory: "Массивы позволяют хранить несколько значений в одной переменной. Элементы массива имеют индексы, начинающиеся с 0.",
        code: `let fruits = ['яблоко', 'банан', 'апельсин'];

console.log(fruits[0]);    // яблоко
console.log(fruits.length); // 3

// Добавление элемента
fruits.push('виноград');
console.log(fruits);       // ['яблоко', 'банан', 'апельсин', 'виноград']`,
        quiz: {
            question: "Как получить длину массива?",
            options: ["array.size", "array.length", "array.count", "length(array)"],
            correct: 1
        }
    },
    {
        id: 5,
        title: "Объекты",
        theory: "Объекты позволяют хранить данные в виде пар ключ-значение. Это основа для создания сложных структур данных.",
        code: `let person = {
    name: 'Иван',
    age: 30,
    city: 'Москва',
    greet: function() {
        return 'Привет, меня зовут ' + this.name;
    }
};

console.log(person.name);    // Иван
console.log(person.greet()); // Привет, меня зовут Иван`,
        quiz: {
            question: "Как получить значение свойства объекта?",
            options: ["object->property", "object.property", "object[property]", "все варианты кроме первого"],
            correct: 3
        }
    },
    {
        id: 6,
        title: "Условные операторы",
        theory: "Условные операторы позволяют выполнять разный код в зависимости от условий. Основные: if, else if, else.",
        code: `let score = 85;

if (score >= 90) {
    console.log('Отлично!');
} else if (score >= 70) {
    console.log('Хорошо!');
} else {
    console.log('Нужно подучить');
}

// Хорошо!`,
        quiz: {
            question: "Что выведет: if (0) console.log('Да'); else console.log('Нет');",
            options: ["Да", "Нет", "Ошибка", "undefined"],
            correct: 1
        }
    },
    {
        id: 7,
        title: "Циклы",
        theory: "Циклы позволяют повторять код множество раз. Основные типы: for, while, do...while, for...of.",
        code: `// Цикл for
for (let i = 0; i < 5; i++) {
    console.log(i); // 0, 1, 2, 3, 4
}

// Цикл while
let count = 0;
while (count < 3) {
    console.log('Итерация', count);
    count++;
}`,
        quiz: {
            question: "Сколько раз выполнится: for (let i = 1; i <= 10; i += 2)",
            options: ["10", "5", "9", "6"],
            correct: 1
        }
    },
    {
        id: 8,
        title: "DOM манипуляции",
        theory: "DOM (Document Object Model) позволяет изменять HTML страницы с помощью JavaScript. Можно добавлять, удалять и изменять элементы.",
        code: `// Найти элемент
let button = document.getElementById('myButton');

// Изменить содержимое
button.textContent = 'Новый текст';

// Добавить обработчик события
button.addEventListener('click', function() {
    alert('Кнопка нажата!');
});`,
        quiz: {
            question: "Как найти элемент по ID?",
            options: ["document.findById()", "document.getElementById()", "document.querySelector()", "второй и третий варианты"],
            correct: 3
        }
    },
    {
        id: 9,
        title: "События",
        theory: "События позволяют реагировать на действия пользователя: клики, нажатия клавиш, загрузку страницы и многое другое.",
        code: `// Обработчик клика
document.addEventListener('click', function(event) {
    console.log('Клик по координатам:', event.clientX, event.clientY);
});

// Обработчик нажатия клавиши
document.addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
        console.log('Нажат Enter');
    }
});`,
        quiz: {
            question: "Как предотвратить действие по умолчанию?",
            options: ["event.stop()", "event.preventDefault()", "event.cancel()", "return false"],
            correct: 1
        }
    },
    {
        id: 10,
        title: "Асинхронность",
        theory: "JavaScript поддерживает асинхронное выполнение кода с помощью промисов, async/await и колбэков. Это позволяет не блокировать выполнение программы.",
        code: `// Промис
function fetchData() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve('Данные получены!');
        }, 1000);
    });
}

// Async/await
async function getData() {
    try {
        let result = await fetchData();
        console.log(result); // Данные получены!
    } catch (error) {
        console.log('Ошибка:', error);
    }
}`,
        quiz: {
            question: "Что используется для работы с асинхронным кодом?",
            options: ["Promise", "async/await", "callbacks", "все варианты"],
            correct: 3
        }
    }
];