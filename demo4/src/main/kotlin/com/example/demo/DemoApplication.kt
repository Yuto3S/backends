package com.example.demo

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.jdbc.core.JdbcTemplate
import org.springframework.stereotype.Service
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController
import java.util.*
import org.springframework.jdbc.core.query
import org.springframework.web.bind.annotation.PathVariable

@SpringBootApplication
class DemoApplication

fun main(args: Array<String>) {
	runApplication<DemoApplication>(*args)
}

@RestController
class MessageController(val service: MessageService) {
	@GetMapping("/")
	fun index(@RequestParam("name") name: String) = "Hello, $name!"

	@GetMapping("/listof")
	fun index2() = listOf(
		Message("1", "Hello"),
		Message("2", "Bonjour"),
		Message("3", "Privet!"),
	)

	@GetMapping("/messages")
	fun findMessages(): List<Message> = service.findMessages()

	@GetMapping("messages/{id}")
	fun getMessageById(@PathVariable id: String): List<Message> = service.findMessageById(id)

	@PostMapping("/messages")
	fun postMessages(@RequestBody message:Message) {
		service.save(message)
	}
}

data class Message(val id: String?, val text: String)

@Service
class MessageService(val db: JdbcTemplate) {
	fun findMessages(): List<Message> = db.query("SELECT * FROM messages") {
		response, _ -> Message(response.getString("id"), response.getString("text"))
	}

	fun findMessageById(id: String): List<Message> = db.query("SELECT * FROM messages WHERE id = ?", id) {
		response, _ -> Message(response.getString("id"), response.getString("text"))
	}

	fun save(message: Message) {
		val id = message.id ?: UUID.randomUUID().toString()
		db.update(
			"INSERT INTO messages VALUES ( ?, ?)",
			id, message.text
		)
	}
}